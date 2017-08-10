// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package log

import (
    "fmt"
    "os"
    "path/filepath"
    "strings"

    "gopkg.in/ini.v1"

    // "github.com/go-stack/stack"
    "github.com/inconshreveable/log15"
    "github.com/inconshreveable/log15/term"
)

var Root log15.Logger
var loggersToClose []DisposableHandler

func init() {
    loggersToClose = make([]DisposableHandler, 0)
    Root = log15.Root()
    Root.SetHandler(log15.DiscardHandler())
}

func New(logger string, ctx ...interface{}) Logger {
    params := append([]interface{}{"logger", logger}, ctx...)
    return Root.New(params...)
}

// Logging
func Trace(format string, v ...interface{}) {
    message := formatLogging(format, v...)
    Root.Debug(message)
}

func Debug(format string, v ...interface{}) {
    message := formatLogging(format, v...)
    Root.Debug(message)
}

func Info(format string, v ...interface{}) {
    message := formatLogging(format, v...)
    Root.Info(message)
}

func Warn(format string, v ...interface{}) {
    message := formatLogging(format, v...)
    Root.Warn(message)
}

func Error(format string, v ...interface{}) {
    message := formatLogging(format, v...)
    Root.Error(message)
}

func Critical(format string, v ...interface{}) {
    message := formatLogging(format, v...)
    Root.Crit(message)
}

func Fatal(format string, v ...interface{}) {
    message := formatLogging(format, v...)
    Root.Crit(message)
    Close()
    os.Exit(1)
}

func formatLogging(format string, v ...interface{}) (string) {
    var message string
    if len(v) > 0 {
        message = fmt.Sprintf(format, v)
    } else {
        message = format
    }
    return message
}

func Close() {
    for _, logger := range loggersToClose {
        logger.Close()
    }
    loggersToClose = make([]DisposableHandler, 0)
}

// func Stack(skip int) string {
//     call := stack.Caller(skip)
//     s := stack.Trace().TrimBelow(call).TrimRuntime()
//     return s.String()
// }

// Setup logging from config
var logLevels = map[string]log15.Lvl{
    "trace":    log15.LvlDebug,
    "debug":    log15.LvlDebug,
    "info":     log15.LvlInfo,
    "warn":     log15.LvlWarn,
    "error":    log15.LvlError,
    "critical": log15.LvlCrit,
}

func getLogLevelFromConfig(key string, defaultName string, cfg *ini.File) (string, log15.Lvl) {
    levelName := cfg.Section(key).Key("level").MustString("warn")
    levelName = strings.ToLower(levelName)
    level := getLogLevelFromString(levelName)
    return levelName, level
}

func getLogLevelFromString(levelName string) log15.Lvl {
    level, ok := logLevels[levelName]

    if !ok {
        Root.Error("Unknown log level", "level", levelName)
        return log15.LvlWarn
    }

    return level
}

func getFilters(filterStrArray []string) map[string]log15.Lvl {
    filterMap := make(map[string]log15.Lvl)

    for _, filterStr := range filterStrArray {
        parts := strings.Split(filterStr, ":")
        if len(parts) > 1 {
            filterMap[parts[0]] = getLogLevelFromString(parts[1])
        }
    }

    return filterMap
}

func getLogFormat(format string) log15.Format {
    switch format {
    case "console":
        if term.IsTty(os.Stdout.Fd()) {
            return log15.TerminalFormat()
        }
        return log15.LogfmtFormat()
    case "text":
        return log15.LogfmtFormat()
    case "json":
        return log15.JsonFormat()
    default:
        return log15.LogfmtFormat()
    }
}

func ReadLoggingConfig(modes []string, logsPath string, cfg *ini.File) {
    Close()

    defaultLevelName, _ := getLogLevelFromConfig("log", "info", cfg)
    defaultFilters := getFilters(cfg.Section("log").Key("filters").Strings(" "))

    handlers := make([]log15.Handler, 0)

    for _, mode := range modes {
        mode = strings.TrimSpace(mode)
        sec, err := cfg.GetSection("log." + mode)
        if err != nil {
            Root.Error("Unknown log mode", "mode", mode)
        }

        // Log level.
        _, level := getLogLevelFromConfig("log."+mode, defaultLevelName, cfg)
        modeFilters := getFilters(sec.Key("filters").Strings(" "))
        format := getLogFormat(sec.Key("format").MustString(""))

        var handler log15.Handler

        // Generate log configuration.
        switch mode {
        case "console":
            handler = log15.StreamHandler(os.Stdout, format)
        case "file":
            fileName := sec.Key("file_name").MustString(filepath.Join(logsPath, "testsite.log"))
            os.MkdirAll(filepath.Dir(fileName), os.ModePerm)
            fileHandler := NewFileWriter()
            fileHandler.Filename = fileName
            fileHandler.Format = format
            fileHandler.Rotate = sec.Key("log_rotate").MustBool(true)
            fileHandler.Maxlines = sec.Key("max_lines").MustInt(1000000)
            fileHandler.Maxsize = 1 << uint(sec.Key("max_size_shift").MustInt(28))
            fileHandler.Daily = sec.Key("daily_rotate").MustBool(true)
            fileHandler.Maxdays = sec.Key("max_days").MustInt64(7)
            fileHandler.Init()

            loggersToClose = append(loggersToClose, fileHandler)
            handler = fileHandler
        case "syslog":
            sysLogHandler := NewSyslog(sec, format)

            loggersToClose = append(loggersToClose, sysLogHandler)
            handler = sysLogHandler
        }

        for key, value := range defaultFilters {
            if _, exist := modeFilters[key]; !exist {
                modeFilters[key] = value
            }
        }

        handler = LogFilterHandler(level, modeFilters, handler)
        handlers = append(handlers, handler)
    }

    Root.SetHandler(log15.MultiHandler(handlers...))
}

func LogFilterHandler(maxLevel log15.Lvl, filters map[string]log15.Lvl, h log15.Handler) log15.Handler {
    return log15.FilterHandler(func(r *log15.Record) (pass bool) {

        if len(filters) > 0 {
            for i := 0; i < len(r.Ctx); i += 2 {
                key := r.Ctx[i].(string)
                if key == "logger" {
                    loggerName, strOk := r.Ctx[i+1].(string)
                    if strOk {
                        if filterLevel, ok := filters[loggerName]; ok {
                            return r.Lvl <= filterLevel
                        }
                    }
                }
            }
        }

        return r.Lvl <= maxLevel
    }, h)
}

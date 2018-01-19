package setting

import (
    "net/url"
    "os"
    "path"
    "path/filepath"
    // "regexp"
    "strings"

    "gopkg.in/ini.v1"

    "github.com/rbak/test-site/pkg/log"
)

func setHomePath() {
    HomePath, _ = filepath.Abs(".")
    // check if homepath is correct
    if pathExists(filepath.Join(HomePath, DEFAULT_CONFIG_PATH)) {
        return
    }

    // try down one path
    if pathExists(filepath.Join(HomePath, "../", DEFAULT_CONFIG_PATH)) {
        HomePath = filepath.Join(HomePath, "../")
    }
}

func pathExists(path string) bool {
    _, err := os.Stat(path)
    if err == nil {
        return true
    }
    if os.IsNotExist(err) {
        return false
    }
    return false
}

func loadConfiguration() {
    var err error

    // load config defaults
    defaultConfigFile := path.Join(HomePath, DEFAULT_CONFIG_PATH)

    Cfg, err = ini.Load(defaultConfigFile)
    if err != nil {
        log.Fatal("Failed to parse defaults.ini, %v", err)
    }

    Cfg.BlockMode = false // speeds up reads

    // evaluate config values containing environment variables
    // evalConfigValues()

    // update data path and logging config
    // DataPath = makeAbsolute(Cfg.Section("paths").Key("data").String(), HomePath)
    initLogging()
}

func initLogging() {
    // split on comma
    LogModes = strings.Split(Cfg.Section("log").Key("mode").MustString("console"), ",")
    // also try space
    if len(LogModes) == 1 {
        LogModes = strings.Split(Cfg.Section("log").Key("mode").MustString("console"), " ")
    }
    LogsPath = makeAbsolute(Cfg.Section("paths").Key("logs").String(), HomePath)
    log.ReadLoggingConfig(LogModes, LogsPath, Cfg)
}

func parseAppUrlAndSubUrl(section *ini.Section) (string, string) {
    appUrl := section.Key("root_url").MustString("http://localhost:3000/")
    if appUrl[len(appUrl)-1] != '/' {
        appUrl += "/"
    }

    // Check if has app suburl.
    url, _ := url.Parse(appUrl)
    // if err != nil {
    //     log.Fatal(4, "Invalid root_url(%s): %s", appUrl, err)
    // }
    appSubUrl := strings.TrimSuffix(url.Path, "/")

    return appUrl, appSubUrl
}

func makeAbsolute(path string, root string) string {
    if filepath.IsAbs(path) {
        return path
    }
    return filepath.Join(root, path)
}

// func evalEnvVarExpression(value string) string {
//     regex := regexp.MustCompile(`\${(\w+)}`)
//     return regex.ReplaceAllStringFunc(value, func(envVar string) string {
//         envVar = strings.TrimPrefix(envVar, "${")
//         envVar = strings.TrimSuffix(envVar, "}")
//         envValue := os.Getenv(envVar)

//         // if env variable is hostname and it is emtpy use os.Hostname as default
//         if envVar == "HOSTNAME" && envValue == "" {
//             envValue, _ = os.Hostname()
//         }

//         return envValue
//     })
// }

// func evalConfigValues() {
//     for _, section := range Cfg.Sections() {
//         for _, key := range section.Keys() {
//             key.SetValue(evalEnvVarExpression(key.Value()))
//         }
//     }
// }

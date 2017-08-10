package setting

import (
    "gopkg.in/ini.v1"

    "github.com/rbak1/test-site/pkg/util"
)

const (
    DEFAULT_CONFIG_PATH string = "conf/defaults.ini"
)

var (
    // Global setting objects.
    Cfg          *ini.File

    // Paths
    HomePath       string
    LogsPath       string
    DataPath       string

    // Log settings.
    LogModes   []string
    LogConfigs []util.DynMap

    // Server settings
    AppUrl         string
    AppSubUrl      string
    HttpAddr       string
    HttpPort       string
    StaticRootPath string
)

func NewConfigContext() error {
    setHomePath()
    loadConfiguration()

    // Server settings
    server := Cfg.Section("server")
    AppUrl, AppSubUrl = parseAppUrlAndSubUrl(server)
    HttpAddr = server.Key("http_addr").MustString("0.0.0.0")
    HttpPort = server.Key("http_port").MustString("3000")
    StaticRootPath = makeAbsolute(server.Key("static_root_path").String(), HomePath)

    return nil
}

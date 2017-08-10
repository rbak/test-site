package main

import (
    "github.com/rbak1/test-site/pkg/log"
    "github.com/rbak1/test-site/pkg/setting"
)

var version = "0.0.1"

func main() {
    server := NewServer()
    server.Start()
}

func initRuntime() {
    err := setting.NewConfigContext()

    if err != nil {
        log.Fatal(err.Error())
    }

    logger := log.New("main")
    logger.Info("Starting server", "version", version)

    // setting.LogConfigurationInfo()
}

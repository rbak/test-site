package main

import (
    "github.com/rbak1/test-site/pkg/log"
    "github.com/rbak1/test-site/pkg/setting"
)

var version = "0.0.1"

func main() {
    log.Init()

    err := setting.NewConfigContext()
    if err != nil {
        log.Fatal(err.Error())
    }

    server := NewServer()
    log.Info("Starting server", "version", version)
    server.Start()
}


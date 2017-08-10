// +build ignore

package main

import (
    "flag"
    "log"
    "os"
    "os/exec"
    "path/filepath"
    "strings"
)

func main() {
    log.SetOutput(os.Stdout)
    log.SetFlags(0)

    ensureGoPath()

    flag.Parse()

    if flag.NArg() == 0 {
        log.Println("Usage: go run build.go build")
        return
    }

    for _, cmd := range flag.Args() {
        switch cmd {
        case "build":
            build("main", "./pkg/main/main.go")
        }
    }
}

func ensureGoPath() {
    if os.Getenv("GOPATH") == "" {
        cwd, err := os.Getwd()
        if err != nil {
            log.Fatal(err)
        }
        gopath := filepath.Clean(filepath.Join(cwd, "../"))
        log.Println("GOPATH is", gopath)
        os.Setenv("GOPATH", gopath)
    }
}

func runPrint(cmd string, args ...string) {
    log.Println(cmd, strings.Join(args, " "))
    ecmd := exec.Command(cmd, args...)
    ecmd.Stdout = os.Stdout
    ecmd.Stderr = os.Stderr
    err := ecmd.Run()
    if err != nil {
        log.Fatal(err)
    }
}

func build(binaryName, pkg string) {
    runPrint("go", "version")
    runPrint("go", "install", "-v", "./pkg/main")
}

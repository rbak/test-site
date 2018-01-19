package main

import (
    "context"
    "fmt"
    "net/http"
    "os"
    "time"

    "github.com/rbak/test-site/pkg/api"
    "github.com/rbak/test-site/pkg/log"
    "github.com/rbak/test-site/pkg/models"
    "github.com/rbak/test-site/pkg/setting"
)

func NewServer() models.Server {
    rootCtx, shutdownFn := context.WithCancel(context.Background())
    // childRoutines, childCtx := errgroup.WithContext(rootCtx)

    return &ServerImpl{
        context:       rootCtx,
        shutdownFn:    shutdownFn,
        log:           log.New("server"),
    }
}

type ServerImpl struct {
    context       context.Context
    shutdownFn    context.CancelFunc
    log           log.Logger
}

func (g *ServerImpl) Start() {
    m := newMacaron()
    api.Register(m)

    listenAddr := fmt.Sprintf("%s:%s", setting.HttpAddr, setting.HttpPort)

    err := http.ListenAndServe(listenAddr, m)

    if err != nil {
        g.Shutdown(1, "Startup failed")
        return
    }
}

func (g *ServerImpl) Shutdown(code int, reason string) {
    g.shutdownFn()
    os.Exit(code)
}

// implement context.Context
func (g *ServerImpl) Deadline() (deadline time.Time, ok bool) {
    return g.context.Deadline()
}
func (g *ServerImpl) Done() <-chan struct{} {
    return g.context.Done()
}
func (g *ServerImpl) Err() error {
    return g.context.Err()
}
func (g *ServerImpl) Value(key interface{}) interface{} {
    return g.context.Value(key)
}


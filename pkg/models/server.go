package models

import "context"

type Server interface {
    context.Context

    Start()
    Shutdown(code int, reason string)
}

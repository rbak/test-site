package api

import (
    "gopkg.in/macaron.v1"
)

func Register(m *macaron.Macaron) {
    // public views
    m.Get("/", Index)
    // m.Get("/login", LoginView)
    // m.Get("/signup", SignupView)
    // m.Get("/architecture", ArchitectureView)
    // m.Get("/logs", LogsView)

    // authed views
    // m.Get("/account", reqSignedIn, AccountView)

    m.NotFound(NotFoundHandler)
}

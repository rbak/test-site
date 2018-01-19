// Copyright 2014 Unknwon
// Copyright 2014 Torkel Ödegaard

package main

import (
    "path"

    "gopkg.in/macaron.v1"

    "github.com/rbak/test-site/pkg/api/static"
    "github.com/rbak/test-site/pkg/middleware"
    "github.com/rbak/test-site/pkg/setting"
)

// var logger log.Logger

func newMacaron() *macaron.Macaron {
    // macaron.Env = setting.Env
    m := macaron.New()

    // m.Use(middleware.Logger())
    // m.Use(middleware.Recovery())

    // if setting.EnableGzip {
    //     m.Use(middleware.Gziper())
    // }

    // for _, route := range plugins.StaticRoutes {
    //     pluginRoute := path.Join("/public/plugins/", route.PluginId)
    //     logger.Debug("Plugins: Adding route", "route", pluginRoute, "dir", route.Directory)
    //     mapStatic(m, route.Directory, "", pluginRoute)
    // }

    mapStatic(m, setting.StaticRootPath, "", "public")
    // mapStatic(m, setting.StaticRootPath, "robots.txt", "robots.txt")
    m.Use(macaron.Renderer(macaron.RenderOptions{
        Directory:  path.Join(setting.StaticRootPath, "views"),
        // IndentJSON: macaron.Env != macaron.PROD,
        // Delims:     macaron.Delims{Left: "[[", Right: "]]"},
    }))

    m.Use(middleware.GetContextHandler())
    // m.Use(middleware.Sessioner(&setting.SessionOptions))
    // m.Use(middleware.RequestMetrics())

    // // needs to be after context handler
    // if setting.EnforceDomain {
    //     m.Use(middleware.ValidateHostHeader(setting.Domain))
    // }

    return m
}

func mapStatic(m *macaron.Macaron, rootDir string, dir string, prefix string) {
    headers := func(c *macaron.Context) {
        c.Resp.Header().Set("Cache-Control", "public, max-age=3600")
    }

    // if setting.Env == setting.DEV {
    //     headers = func(c *macaron.Context) {
    //         c.Resp.Header().Set("Cache-Control", "max-age=0, must-revalidate, no-cache")
    //     }
    // }

    m.Use(httpstatic.Static(
        path.Join(rootDir, dir),
        httpstatic.StaticOptions{
            SkipLogging: true,
            Prefix:      prefix,
            AddHeaders:  headers,
        },
    ))
}

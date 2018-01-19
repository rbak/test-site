package api

import (
    "github.com/rbak/test-site/pkg/api/dtos"
    "github.com/rbak/test-site/pkg/middleware"
    "github.com/rbak/test-site/pkg/setting"
)

func setIndexViewData(c *middleware.Context) (*dtos.IndexViewData, error) {
    // settings, err := getFrontendSettingsMap(c)
    // if err != nil {
    //     return nil, err
    // }

    // prefsQuery := m.GetPreferencesWithDefaultsQuery{OrgId: c.OrgId, UserId: c.UserId}
    // if err := bus.Dispatch(&prefsQuery); err != nil {
    //     return nil, err
    // }
    // prefs := prefsQuery.Result

    appUrl := setting.AppUrl
    appSubUrl := setting.AppSubUrl

    var data = dtos.IndexViewData{
        // User: &dtos.CurrentUser{
        //     Id:             c.UserId,
        //     IsSignedIn:     c.IsSignedIn,
        //     Login:          c.Login,
        //     LightTheme:     prefs.Theme == "light",
        // },
        // Settings:                settings,
        AppUrl:                  appUrl,
        AppSubUrl:               appSubUrl,
        // BuildVersion:            setting.BuildVersion,
        // BuildCommit:             setting.BuildCommit,
    }

    // if len(data.User.Name) == 0 {
    //     data.User.Name = data.User.Login
    // }

    // themeUrlParam := c.Query("theme")
    // if themeUrlParam == "light" {
    //     data.User.LightTheme = true
    // }

    return &data, nil
}

func Index(c *middleware.Context) {
    if data, err := setIndexViewData(c); err != nil {
        c.Handle(500, "Failed to get settings", err)
        return
    } else {
        c.HTML(200, "index", data)
    }
}

func NotFoundHandler(c *middleware.Context) {
    if data, err := setIndexViewData(c); err != nil {
        c.Handle(500, "Failed to get settings", err)
        return
    } else {
        c.HTML(404, "index", data)
    }
}

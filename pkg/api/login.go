package api

// func LoginView(c *middleware.Context) {
//     viewData, err := setIndexViewData(c)
//     if err != nil {
//         c.Handle(500, "Failed to get settings", err)
//         return
//     }

//     enabledOAuths := make(map[string]interface{})
//     for key, oauth := range setting.OAuthService.OAuthInfos {
//         enabledOAuths[key] = map[string]string{"name": oauth.Name}
//     }

//     viewData.Settings["oauth"] = enabledOAuths
//     viewData.Settings["disableUserSignUp"] = !setting.AllowUserSignUp
//     viewData.Settings["loginHint"] = setting.LoginHint
//     viewData.Settings["disableLoginForm"] = setting.DisableLoginForm

//     if !tryLoginUsingRememberCookie(c) {
//         c.HTML(200, VIEW_INDEX, viewData)
//         return
//     }

//     if redirectTo, _ := url.QueryUnescape(c.GetCookie("redirect_to")); len(redirectTo) > 0 {
//         c.SetCookie("redirect_to", "", -1, setting.AppSubUrl+"/")
//         c.Redirect(redirectTo)
//         return
//     }

//     c.Redirect(setting.AppSubUrl + "/")
// }

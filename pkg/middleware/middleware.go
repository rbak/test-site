package middleware

import (
    "strconv"

    "gopkg.in/macaron.v1"

    "github.com/rbak/test-site/pkg/setting"
)

type Context struct {
    *macaron.Context
    // *m.SignedInUser

    // Session SessionStore

    // IsSignedIn     bool
    // Logger         log.Logger
}

func GetContextHandler() macaron.Handler {
    return func(c *macaron.Context) {
        ctx := &Context{
            Context:        c,
            // SignedInUser:   &m.SignedInUser{},
            // Session:        GetSession(),
            // IsSignedIn:     false,
            // Logger:         log.New("context"),
        }

        // the order in which these are tested are important
        // look for api key in Authorization header first
        // then init session and look for userId in session
        // then look for api key in session (special case for render calls via api)
        // then test if anonymous access is enabled
        // if initContextWithBasicAuth(ctx) ||
        //     initContextWithAuthProxy(ctx) ||
        //     initContextWithUserSessionCookie(ctx) ||
        //     initContextWithAnonymousUser(ctx) {
        // }

        // ctx.Logger = log.New("context", "userId", ctx.UserId, "orgId", ctx.OrgId, "uname", ctx.Login)
        ctx.Data["ctx"] = ctx

        c.Map(ctx)
    }
}

// func initContextWithUserSessionCookie(ctx *Context) bool {
//     // initialize session
//     if err := ctx.Session.Start(ctx); err != nil {
//         ctx.Logger.Error("Failed to start session", "error", err)
//         return false
//     }

//     var userId int64
//     if userId = getRequestUserId(ctx); userId == 0 {
//         return false
//     }

//     query := m.GetSignedInUserQuery{UserId: userId}
//     if err := bus.Dispatch(&query); err != nil {
//         ctx.Logger.Error("Failed to get user with id", "userId", userId)
//         return false
//     } else {
//         ctx.SignedInUser = query.Result
//         ctx.IsSignedIn = true
//         return true
//     }
// }

// func initContextWithBasicAuth(ctx *Context) bool {
//     if !setting.BasicAuthEnabled {
//         return false
//     }

//     header := ctx.Req.Header.Get("Authorization")
//     if header == "" {
//         return false
//     }

//     username, password, err := util.DecodeBasicAuthHeader(header)
//     if err != nil {
//         ctx.JsonApiErr(401, "Invalid Basic Auth Header", err)
//         return true
//     }

//     loginQuery := m.GetUserByLoginQuery{LoginOrEmail: username}
//     if err := bus.Dispatch(&loginQuery); err != nil {
//         ctx.JsonApiErr(401, "Basic auth failed", err)
//         return true
//     }

//     user := loginQuery.Result

//     // validate password
//     if util.EncodePassword(password, user.Salt) != user.Password {
//         ctx.JsonApiErr(401, "Invalid username or password", nil)
//         return true
//     }

//     query := m.GetSignedInUserQuery{UserId: user.Id}
//     if err := bus.Dispatch(&query); err != nil {
//         ctx.JsonApiErr(401, "Authentication error", err)
//         return true
//     } else {
//         ctx.SignedInUser = query.Result
//         ctx.IsSignedIn = true
//         return true
//     }
// }

// Handle handles and logs error by given status.
func (ctx *Context) Handle(status int, title string, err error) {
    // if err != nil {
    //     ctx.Logger.Error(title, "error", err)
    //     if setting.Env != setting.PROD {
    //         ctx.Data["ErrorMsg"] = err
    //     }
    // }

    ctx.Data["Title"] = title
    ctx.Data["AppSubUrl"] = setting.AppSubUrl
    ctx.HTML(status, strconv.Itoa(status))
}

// func (ctx *Context) JsonOK(message string) {
//     resp := make(map[string]interface{})
//     resp["message"] = message
//     ctx.JSON(200, resp)
// }

// func (ctx *Context) HasUserRole(role m.RoleType) bool {
//     return ctx.OrgRole.Includes(role)
// }

// func (ctx *Context) TimeRequest(timer metrics.Timer) {
//     ctx.Data["perfmon.timer"] = timer
// }

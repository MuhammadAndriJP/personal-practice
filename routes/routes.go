package routes

import (
	// "project_altabe4_1/external/google"

	"github.com/MuhAndriJP/personal-practice.git/external/google"
	"github.com/MuhAndriJP/personal-practice.git/external/xendit"
	"github.com/MuhAndriJP/personal-practice.git/handler/user"
	"github.com/MuhAndriJP/personal-practice.git/web"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	e.Use(middleware.RequestID())
	e.Use(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(middleware.Gzip())
	e.Use(middleware.Static("."))
	e.Use(middleware.CORS())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization, echo.HeaderContentLength, echo.HeaderAcceptEncoding, echo.HeaderAccessControlAllowOrigin, echo.HeaderAccessControlAllowHeaders, echo.HeaderContentDisposition, "X-Request-Id", "device-id", "X-Summary", "X-Account-Number", "X-Business-Name", "client-secret", "X-CSRF-Token", "app-key"},
		ExposeHeaders:    []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization, echo.HeaderContentLength, echo.HeaderAcceptEncoding, echo.HeaderAccessControlAllowOrigin, echo.HeaderAccessControlAllowHeaders, echo.HeaderContentDisposition, "X-Request-Id", "device-id", "X-Summary", "X-Account-Number", "X-Business-Name", "client-secret", "X-CSRF-Token", "app-key"},
		AllowMethods:     []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
		AllowCredentials: true,
	}))

	// WEB
	e.GET("/index", web.IndexHandler)
	e.GET("/register", web.RegisterHandler)
	e.GET("/login", web.LoginHandler)

	g := e.Group("api")
	// Login And Register
	g.POST("/users/register", user.NewUserRegister().Handle)
	g.POST("/users/login", user.NewUserLogin().Handle)

	// Google
	g.GET("/google", google.NewGoogleAuth().HandleGoogleLogin)
	g.GET("/UserOauth", google.NewGoogleAuth().HandleGoogleCallback)

	// Xendit
	g.POST("/xendit/ewallet/charge", xendit.CreateEWalletCharge)
	g.GET("/xendit/ewallet/status/:id", xendit.GetEWalletChargeStatus)
	g.POST("/xendit/ewallet/callback", xendit.CreateEWalletCallback)

	return e
}

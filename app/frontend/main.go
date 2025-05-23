// Code generated by hertz generator.

package main

import (
	"context"
	"fmt" // For error handling during init
	"os"
	"time"

	"gomall/app/frontend/biz/router"
	"gomall/app/frontend/conf"
	"gomall/app/frontend/infra/rpc"
	"gomall/app/frontend/middleware"
	frontendUtils "gomall/app/frontend/utils" // Ensure this points to your frontend utils

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/cors"
	"github.com/hertz-contrib/gzip"
	"github.com/hertz-contrib/logger/accesslog"
	hertzlogrus "github.com/hertz-contrib/logger/logrus"
	"github.com/hertz-contrib/pprof"
	"github.com/hertz-contrib/sessions"
	"github.com/hertz-contrib/sessions/redis"
	"github.com/joho/godotenv"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	// init dal
	// dal.Init()
	err := godotenv.Load() // Loads .env from the current directory (app/frontend)
	if err != nil {
		panic(fmt.Sprintf("Error loading .env file in frontend service: %v", err))
	}
	frontendUtils.InitJWTKey()            // Initialize JWT secret key for frontend
	rpc.Init()
	address := conf.GetConf().Hertz.Address
	h := server.New(server.WithHostPorts(address))

	registerMiddleware(h)

	// add a ping route to test
	h.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(consts.StatusOK, utils.H{"ping": "pong"})
	})

	router.GeneratedRegister(h)
	h.LoadHTMLGlob("template/*")
	h.Static("/static", "./")
	h.GET("/sign-in", func(c context.Context, ctx *app.RequestContext) {
		data := utils.H{
			"Title": "Sign In",
			"Next":  ctx.Query("next"),
		}
		ctx.HTML(consts.StatusOK, "sign-in", data)
	})
	h.GET("/sign-up", func(c context.Context, ctx *app.RequestContext) {
		ctx.HTML(consts.StatusOK, "sign-up", utils.H{"Title": "Sign Up"})
	})
	h.GET("/about", middleware.Auth(), func(c context.Context, ctx *app.RequestContext) {
		ctx.HTML(consts.StatusOK, "about", utils.H{"Title": "About"})
	})
	h.Spin()
}

func registerMiddleware(h *server.Hertz) {
	// log
	store, _ := redis.NewStore(10, "tcp", conf.GetConf().Redis.Address, "", []byte(os.Getenv("SESSION_SECRET")))
	h.Use(sessions.New("mcc-shop", store))

	logger := hertzlogrus.NewLogger()
	hlog.SetLogger(logger)
	hlog.SetLevel(conf.LogLevel())
	asyncWriter := &zapcore.BufferedWriteSyncer{
		WS: zapcore.AddSync(&lumberjack.Logger{
			Filename:   conf.GetConf().Hertz.LogFileName,
			MaxSize:    conf.GetConf().Hertz.LogMaxSize,
			MaxBackups: conf.GetConf().Hertz.LogMaxBackups,
			MaxAge:     conf.GetConf().Hertz.LogMaxAge,
		}),
		FlushInterval: time.Minute,
	}
	hlog.SetOutput(asyncWriter)
	h.OnShutdown = append(h.OnShutdown, func(ctx context.Context) {
		asyncWriter.Sync()
	})

	// pprof
	if conf.GetConf().Hertz.EnablePprof {
		pprof.Register(h)
	}

	// gzip
	if conf.GetConf().Hertz.EnableGzip {
		h.Use(gzip.Gzip(gzip.DefaultCompression))
	}

	// access log
	if conf.GetConf().Hertz.EnableAccessLog {
		h.Use(accesslog.New())
	}

	// recovery
	h.Use(recovery.Recovery())

	// cores
	h.Use(cors.Default())
	//middlerware
	middleware.Register(h)
}

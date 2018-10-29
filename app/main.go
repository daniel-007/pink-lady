// gin-skeleton: Typically gin-based web application's organizational structure
package main

import (
	"fmt"
	"os"
	"strings"
	"syscall"

	"github.com/axiaoxin/gin-skeleton/app/apis"
	"github.com/axiaoxin/gin-skeleton/app/common"
	"github.com/axiaoxin/gin-skeleton/app/middleware"
	"github.com/axiaoxin/gin-skeleton/app/utils"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func main() {
	defer utils.DB.Close()
	// TODO: imp in cli
	version := pflag.Bool("version", false, "show version")
	check := pflag.Bool("check", false, "check everything need to be checked")
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
	if *version {
		fmt.Println(common.VERSION)
		os.Exit(0)
	}
	if *check {
		fmt.Println("I'm fine :)")
		os.Exit(0)
	}
	mode := strings.ToLower(viper.GetString("server.mode"))
	if mode == "debug" {
		gin.SetMode(gin.DebugMode)
		utils.DB.LogMode(true)
	} else if mode == "test" {
		utils.DB.LogMode(true)
		gin.SetMode(gin.TestMode)
	} else {
		utils.DB.LogMode(false)
		gin.DisableConsoleColor()
		gin.SetMode(gin.ReleaseMode)
	}

	app := gin.Default()
	app.Use(middleware.RequestID())

	apis.RegisterRoutes(app)

	server := endless.NewServer(viper.GetString("server.bind"), app)
	server.BeforeBegin = func(addr string) {
		logrus.Infof("Gin server is listening and serving HTTP on %s (pids: %d)", addr, syscall.Getpid())
	}
	server.ListenAndServe()
}

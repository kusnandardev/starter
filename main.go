package main

import (
	"kusnandartoni/starter/models"
	"kusnandartoni/starter/pkg/logging"
	"kusnandartoni/starter/pkg/setting"
	"kusnandartoni/starter/redisdb"
	"kusnandartoni/starter/routers"
	"fmt"
	"log"
	"net/http"
	"time"
)

func init() {
	now := time.Now()
	setting.Setup()
	logging.Setup()
	models.Setup()
	redisdb.Setup()
	log.Printf("All setup is done in %v \n", time.Since(now))
}

// @title Starter
// @version 1.0
// @description Backend REST API for golang starter

// @contact.name Toni Kusnandar
// @contact.url https://www.linkedin.com/in/kusnandartoni/
// @contact.email kusnandartoni@gmail.com

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	routersInit := routers.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HTTPPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	logging.Info("0", "start http server listening "+endPoint)

	server.ListenAndServe()
}

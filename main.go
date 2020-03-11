package main

import (
	"fmt"
	"kusnandartoni/starter/models"
	"kusnandartoni/starter/pkg/logging"
	"kusnandartoni/starter/pkg/pool"
	"kusnandartoni/starter/pkg/setting"
	"kusnandartoni/starter/redisdb"
	"kusnandartoni/starter/routers"
	"log"
	"net/http"
	"time"

	"github.com/jasonlvhit/gocron"
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

	if setting.ServerSetting.RunMode != "debug" {
		logging.Info("start http server listening " + endPoint)
	} else {
		fmt.Println("start http server listening " + endPoint)
	}
	go runCrond()
	server.ListenAndServe()

}

func runCrond() {
	gocron.Every(10).Seconds().Do(sendMail)
	<-gocron.Start()
}

func sendMail() {
	var data []string
	key := "starter_email"

	// tStart := time.Now()
	// log.Printf("Starting Application at %s\n", tStart.Format("2006-01-02 15:04:05"))

	data, err := redisdb.GetList(key)
	if err != nil {
		log.Fatal(err)
	}
	dataLen := len(data)
	if dataLen > 100 {
		data = data[0:100]
	}
	// log.Println(data, "len: ", len(data))
	if len(data) > 0 {
		err = redisdb.RemoveList(key, data)
		if err != nil {
			log.Fatal(err)
		}
	}
	// wg.Add(len(data))
	collector := pool.StartDispatcher(5)
	for i, job := range data {
		collector.Work <- pool.Work{Job: job, ID: i}
	}

	// wg.Wait()
	// tStop := time.Now()
	// diff := tStop.Sub(tStart)
	// log.Printf("Application Stoped at %s", tStop.Format("2006-01-02 15:04:05"))
	// log.Printf("Application running for %v \n\n", diff)
}

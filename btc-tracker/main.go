package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/emicklei/go-restful"
	"github.com/pho-coder/ants-plan/btc-tracker/apps"
	"github.com/pho-coder/ants-plan/btc-tracker/config"
)

//HealthCheck ...
func HealthCheck(req *restful.Request, resp *restful.Response) {
	log.Println("come on Healthcheck from " + req.SelectedRoutePath())
	io.WriteString(resp, fmt.Sprintf("%s", time.Now().Format(config.TIME_FORMAT)))
}

// CurrentPrice ...
func CurrentPrice(req *restful.Request, resp *restful.Response) {
	log.Println("come on CurrentPrice from " + req.SelectedRoutePath())
	currentPrice, _ := json.Marshal(apps.GetCurrentPrice())
	io.WriteString(resp, string(currentPrice))
}

// CurrentData ...
func CurrentData(req *restful.Request, resp *restful.Response) {
	log.Println("come on CurrentData from " + req.SelectedRoutePath())
	io.WriteString(resp, string("test"))
}

func index(req *restful.Request, resp *restful.Response) {
	log.Println("come on index from " + req.SelectedRoutePath())
	io.WriteString(resp, "<a href='/healthcheck'>HealthCheck</a><br>"+"<a href='/currentprice'>CurrentPrice</a><br>"+
		"<a href='/currentdata'>CurrentData</a><br>")
}

func main() {
	webservice := new(restful.WebService)
	webservice.Route(webservice.GET("/").To(index))
	webservice.Route(webservice.GET("/healthcheck").To(HealthCheck))
	webservice.Route(webservice.GET("/currentprice").To(CurrentPrice))
	webservice.Route(webservice.GET("/currentdata").To(CurrentData))
	restful.Add(webservice)
	srv := &http.Server{
		Addr:         "0.0.0.0:7001",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}

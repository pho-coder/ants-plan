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
)

//HealthCheck ...
func HealthCheck(req *restful.Request, resp *restful.Response) {
	log.Println("come on Healthcheck from " + req.SelectedRoutePath())
	io.WriteString(resp, fmt.Sprintf("%s", time.Now()))
}

// CurrentPrice ...
func CurrentPrice(req *restful.Request, resp *restful.Response) {
	log.Println("come on CurrentPrice from " + req.SelectedRoutePath())
	currentPrice, _ := json.Marshal(apps.GetCurrentPrice())
	io.WriteString(resp, string(currentPrice))
}

func index(req *restful.Request, resp *restful.Response) {
	log.Println("come on index from " + req.SelectedRoutePath())
	io.WriteString(resp, "<a href='/healthcheck'>HealthCheck</a>\n"+"<a href='/currentprice'>CurrentPrice</a>")
}

func main() {
	webservice := new(restful.WebService)
	webservice.Route(webservice.GET("/").To(index))
	webservice.Route(webservice.GET("/healthcheck").To(HealthCheck))
	webservice.Route(webservice.GET("/currentprice").To(CurrentPrice))
	restful.Add(webservice)
	srv := &http.Server{
		Addr:         "0.0.0.0:7001",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}

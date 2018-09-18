package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/emicklei/go-restful"
	"github.com/pho-coder/ants-plan/btc-tracker/apps"
)

//Healthcheck ...
func Healthcheck(req *restful.Request, resp *restful.Response) {
	log.Println("come on Healthcheck from " + req.SelectedRoutePath())
	io.WriteString(resp, fmt.Sprintf("%s", time.Now()))
}

// CurrentPrice ...
func CurrentPrice(req *restful.Request, resp *restful.Response) {
	log.Println("come on CurrentPrice from " + req.SelectedRoutePath())
	io.WriteString(resp, strconv.FormatFloat(apps.GetCurrentPrice(), 'f', -1, 64))
}

func main() {

	webservice := new(restful.WebService)
	webservice.Route(webservice.GET("/healthcheck").To(Healthcheck))
	webservice.Route(webservice.GET("/").To(Healthcheck))
	webservice.Route(webservice.GET("/currentprice").To(CurrentPrice))
	restful.Add(webservice)
	srv := &http.Server{
		Addr:         "0.0.0.0:7001",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}

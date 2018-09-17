package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/emicklei/go-restful"
)

//Healthcheck ...
func Healthcheck(req *restful.Request, resp *restful.Response) {
	log.Println("come on Healthcheck from " + req.SelectedRoutePath())
	io.WriteString(resp, fmt.Sprintf("%s", time.Now()))
}

//

func main() {

	webservice := new(restful.WebService)
	webservice.Route(webservice.GET("/healthcheck").To(Healthcheck))
	webservice.Route(webservice.GET("/").To(Healthcheck))
	restful.Add(webservice)
	srv := &http.Server{
		Addr:         "0.0.0.0:7001",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}

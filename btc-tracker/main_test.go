package main_test

import (
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/pho-coder/ants-plan/btc-tracker/config"
)

func TestGetHealthCheckURL(t *testing.T) {
	// make a dummy request
	response, err := http.Get("http://localhost:7001/healthcheck")
	if http.StatusOK != response.StatusCode {
		t.Errorf("Expected response code %d. Got %d\n", http.StatusOK, response.StatusCode)
	}
	if err != nil {
		t.Errorf("Encountered an error: %s", err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Errorf("Encountered an error: %s", err)
	}
	timeGot, _ := time.Parse(config.TIME_FORMAT, string(body))
	subM := time.Now().Sub(timeGot).Minutes()
	if subM >= -10 && subM <= 10 {
		t.Log("healthcheck time is OK")
	} else {
		t.Errorf("healthcheck time is over: %f", subM)
	}
}

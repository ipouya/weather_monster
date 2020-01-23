package http_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"weather-monster/models"
	"weather-monster/testHelper"
)

func Test_PostTemperature(t *testing.T) {
	testRouter := testHelper.StartTestServer()
	city := testHelper.MakeCity()
	data,err := json.Marshal(&models.Temperature{Max:35,Min:30,CityId:city.ID})
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	req, _ := http.NewRequest("POST", "/temperatures", bytes.NewBuffer(data))
	w := httptest.NewRecorder()

	testRouter.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Status code should be %v, was %d.", http.StatusOK, w.Code)
		return
	}

	decoder := json.NewDecoder(w.Body)
	var body models.Temperature
	err = decoder.Decode(&body)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	if body.ID == 0 {
		t.Errorf("ID must not be zero")
	}
}

func Test_Forecasts(t *testing.T) {
	testRouter := testHelper.StartTestServer()
	city := testHelper.MakeCity()
	sumMax := 0
	sumMin := 0
	count := 100
	for i:=0 ; i<count ;i++ {
		temp := testHelper.MakeTemperature(city.ID)
		sumMax += temp.Max
		sumMin += temp.Min
	}
	url := fmt.Sprint("/forecasts/",city.ID)
	req, _ := http.NewRequest("GET", url, nil)
	w := httptest.NewRecorder()

	testRouter.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Status code should be %v, was %d.", http.StatusOK, w.Code)
		return
	}

	decoder := json.NewDecoder(w.Body)
	var body models.Forecast
	err := decoder.Decode(&body)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	if body.Min != sumMin/count {
		t.Errorf("Min-Average should be %d, was %d.",sumMin/count , body.Min)
	}
	if body.Max != sumMax/count {
		t.Errorf("Max-Average should be %d, was %d.",sumMax/count , body.Max)
	}
	if body.Sample != count {
		t.Errorf("Sample should be %d, was %d.",count , body.Sample)
	}
}

package http_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	"weather-monster/models"
	"weather-monster/testHelper"
)

func Test_PostCity(t *testing.T) {
	testRouter := testHelper.StartTestServer()
	city := &models.City{
		Name: fmt.Sprint(time.Now().Unix()),
		Latitude: (rand.Float64() * 5) + 5,
		Longitude: (rand.Float64() * 5) + 5,
	}
	data,err := json.Marshal(city)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	req, _ := http.NewRequest("POST", "/cities", bytes.NewBuffer(data))
	w := httptest.NewRecorder()

	testRouter.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Status code should be %v, was %d.", http.StatusOK, w.Code)
		return
	}

	decoder := json.NewDecoder(w.Body)
	var body models.City
	err = decoder.Decode(&body)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	if body.ID == 0 {
		t.Errorf("ID must not be zero")
	}
}


package http_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"weather-monster/models"
	"weather-monster/testHelper"
)

func Test_PostWebhook(t *testing.T) {
	testRouter := testHelper.StartTestServer()
	city := testHelper.MakeCity()
	data,err := json.Marshal(&models.Webhook{CityId:city.ID,CallbackUrl:"http://localhost:3000/test"})
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	req, _ := http.NewRequest("POST", "/webhooks", bytes.NewBuffer(data))
	w := httptest.NewRecorder()

	testRouter.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Status code should be %v, was %d.", http.StatusOK, w.Code)
		return
	}

	decoder := json.NewDecoder(w.Body)
	var body models.Webhook
	err = decoder.Decode(&body)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	if body.ID == 0 {
		t.Errorf("ID must not be zero")
	}
}


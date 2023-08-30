package webapp

import (
	"Goondex/part13/pkg/crawler"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAPI_docs(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/v1/docs", nil)
	rec := httptest.NewRecorder()
	api.router.ServeHTTP(rec, req)
	if !(rec.Code == http.StatusOK) {
		t.Errorf("код неверен: получили %d, а хотели %d", rec.Code, http.StatusOK)
	}
	t.Log("Response: ", rec.Body)

	var body []crawler.Document
	err := json.NewDecoder(rec.Body).Decode(&body)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	want := 2
	got := len(body)

	if got != want {
		t.Errorf("ответ неверен: получили %v элемента, а хотели %v элемента", got, want)
	}
}

func TestAPI_doc(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/v1/docs/1", nil)
	rec := httptest.NewRecorder()
	api.router.ServeHTTP(rec, req)
	if !(rec.Code == http.StatusOK) {
		t.Errorf("код неверен: получили %d, а хотели %d", rec.Code, http.StatusOK)
	}

	//----------------------------------------------------------------

	var body crawler.Document
	err := json.NewDecoder(rec.Body).Decode(&body)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	want := "1 URL2 Title2"
	got := fmt.Sprintf("%v %v %v", body.ID, body.URL, body.Title)

	if got != want {
		t.Errorf("ответ неверен: получили %v, а хотели %v", got, want)
	}

	//----------------------------------------------------------------

	req = httptest.NewRequest(http.MethodGet, "/api/v1/docs/3", nil)
	rec = httptest.NewRecorder()
	api.router.ServeHTTP(rec, req)

	if !(rec.Code == http.StatusNotFound) {
		t.Errorf("код неверен: получили %d, а хотели %d", rec.Code, http.StatusNotFound)
	}

	//----------------------------------------------------------------

	req = httptest.NewRequest(http.MethodGet, "/api/v1/docs/not_exist", nil)
	rec = httptest.NewRecorder()
	api.router.ServeHTTP(rec, req)

	if !(rec.Code == http.StatusUnprocessableEntity) {
		t.Errorf("код неверен: получили %d, а хотели %d", rec.Code, http.StatusUnprocessableEntity)
	}

	want = "Некорректный id документа: not_exist\n"
	data, err := ioutil.ReadAll(rec.Body)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	got = string(data)

	if got != want {
		t.Errorf("ответ неверен: получили %v, а хотели %v", got, want)
	}
}

func TestAPI_newDoc(t *testing.T) {
	doc := crawler.Document{
		ID:    2,
		URL:   "http://basistech.ru",
		Title: "BASIS",
		Body:  "",
	}
	json_doc, _ := json.Marshal(doc)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/docs", bytes.NewBuffer(json_doc))
	rec := httptest.NewRecorder()
	api.router.ServeHTTP(rec, req)
	if !(rec.Code == http.StatusCreated) {
		t.Errorf("код неверен: получили %d, а хотели %d", rec.Code, http.StatusCreated)
	}

	//----------------------------------------------------------------

	var body crawler.Document
	err := json.NewDecoder(rec.Body).Decode(&body)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	want := "2 http://basistech.ru BASIS"
	got := fmt.Sprintf("%v %v %v", body.ID, body.URL, body.Title)

	if got != want {
		t.Errorf("ответ неверен: получили %v, а хотели %v", got, want)
	}

	//----------------------------------------------------------------

	invalid_doc := []byte(`{"Any":"Other","Keys":"Send"}`)
	json_doc, _ = json.Marshal(invalid_doc)
	req = httptest.NewRequest(http.MethodPost, "/api/v1/docs", bytes.NewBuffer(json_doc))
	rec = httptest.NewRecorder()
	api.router.ServeHTTP(rec, req)
	if !(rec.Code == http.StatusUnprocessableEntity) {
		t.Errorf("код неверен: получили %d, а хотели %d", rec.Code, http.StatusUnprocessableEntity)
	}
}

func TestAPI_delDoc(t *testing.T) {
	req := httptest.NewRequest(http.MethodDelete, "/api/v1/docs/1", nil)
	rec := httptest.NewRecorder()
	api.router.ServeHTTP(rec, req)
	if !(rec.Code == http.StatusOK) {
		t.Errorf("код неверен: получили %d, а хотели %d", rec.Code, http.StatusOK)
	}

	//----------------------------------------------------------------

	want := ""
	data, err := ioutil.ReadAll(rec.Body)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	got := string(data)

	if got != want {
		t.Errorf("ответ неверен: получили %v, а хотели %v", got, want)
	}

	//----------------------------------------------------------------

	req = httptest.NewRequest(http.MethodDelete, "/api/v1/docs/3", nil)
	rec = httptest.NewRecorder()
	api.router.ServeHTTP(rec, req)
	if !(rec.Code == http.StatusNotFound) {
		t.Errorf("код неверен: получили %d, а хотели %d", rec.Code, http.StatusOK)
	}
}

func TestAPI_updateDoc(t *testing.T) {
	doc := crawler.Document{
		ID:    0,
		URL:   "http://basistech.ru",
		Title: "BASIS",
		Body:  "",
	}
	json_doc, _ := json.Marshal(doc)
	req := httptest.NewRequest(http.MethodPut, "/api/v1/docs/0", bytes.NewBuffer(json_doc))
	rec := httptest.NewRecorder()
	api.router.ServeHTTP(rec, req)
	if !(rec.Code == http.StatusOK) {
		t.Errorf("код неверен: получили %d, а хотели %d", rec.Code, http.StatusOK)
	}

	//----------------------------------------------------------------

	var body crawler.Document
	err := json.NewDecoder(rec.Body).Decode(&body)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	want := "0 http://basistech.ru BASIS"
	got := fmt.Sprintf("%v %v %v", body.ID, body.URL, body.Title)

	if got != want {
		t.Errorf("ответ неверен: получили %v, а хотели %v", got, want)
	}

	//----------------------------------------------------------------

	invalid_doc := []byte(`{"Any":"Other","Keys":"Send"}`)
	json_doc, _ = json.Marshal(invalid_doc)
	req = httptest.NewRequest(http.MethodPost, "/api/v1/docs", bytes.NewBuffer(json_doc))
	rec = httptest.NewRecorder()
	api.router.ServeHTTP(rec, req)
	if !(rec.Code == http.StatusUnprocessableEntity) {
		t.Errorf("код неверен: получили %d, а хотели %d", rec.Code, http.StatusUnprocessableEntity)
	}
}

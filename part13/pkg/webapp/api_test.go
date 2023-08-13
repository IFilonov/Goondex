package webapp

import (
	"Goondex/part13/pkg/crawler"
	"os"
	"testing"

	"github.com/gorilla/mux"
)

var api *API

func TestMain(m *testing.M) {
	apiData.Docs = []crawler.Document{
		{ID: 0, URL: "URL1", Title: "Title1", Body: ""},
		{ID: 1, URL: "URL2", Title: "Title2", Body: ""}}

	api = new(API)
	api.router = mux.NewRouter()
	api.Endpoints()
	os.Exit(m.Run())
}

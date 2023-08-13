package webapp

import (
	"Goondex/part13/pkg/crawler"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ApiData struct {
	IndexDocs []byte
	Docs      []crawler.Document
}

func (api *API) docs(w http.ResponseWriter, r *http.Request) {

	json_docs, err := json.Marshal(apiData.Docs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("Ошибка сериализации списка документов")
		return
	}

	_, err = w.Write(json_docs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (api *API) newDoc(w http.ResponseWriter, r *http.Request) {
	var doc crawler.Document

	err := json.NewDecoder(r.Body).Decode(&doc)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	lastDoc := apiData.Docs[len(apiData.Docs)-1]
	doc.ID = lastDoc.ID + 1
	apiData.Docs = append(apiData.Docs, doc)

	json_doc, err := json.Marshal(doc)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("Ошибка сериализации списка документов")
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, err = w.Write(json_doc)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (api *API) doc(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, fmt.Sprintf("Некорректный id документа: %v", vars["id"]), http.StatusUnprocessableEntity)
		return
	}

	l := len(apiData.Docs)
	if id < 0 || id > l {
		http.Error(w, fmt.Sprintf("Не найден документ с id: %v", id), http.StatusNotFound)
		return
	}

	json_doc, err := json.Marshal(apiData.Docs[id])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("Ошибка сериализации списка документов")
		return
	}

	_, err = w.Write(json_doc)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (api *API) editDoc(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		fmt.Println("Ошибка передачи id документа")
		return
	}

	l := len(apiData.Docs)
	if id < 0 || id > l {
		http.Error(w, fmt.Sprintf("Не найден документ с id %v", id), http.StatusNotFound)
		fmt.Println("Не найден документ")
		return
	}

	var new_doc crawler.Document
	err = json.NewDecoder(r.Body).Decode(&new_doc)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		fmt.Printf("error: %v", err)
		return
	}

	apiData.Docs[id] = new_doc

	json_doc, err := json.Marshal(apiData.Docs[id])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("Ошибка сериализации списка документов")
		return
	}

	_, err = w.Write(json_doc)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (api *API) deleteDoc(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		fmt.Println("Ошибка передачи id документа")
		return
	}

	l := len(apiData.Docs)
	if id < 0 || id > l {
		http.Error(w, fmt.Sprintf("Не найден документ с id %v", id), http.StatusNotFound)
		fmt.Println("Не найден документ")
		return
	}

	apiData.Docs = append(apiData.Docs[:id], apiData.Docs[id+1:]...)

	_, err = w.Write(nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

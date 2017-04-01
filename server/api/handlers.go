package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// Document : data model for a document object
// TODO: Move to models directory
type Document struct {
	ID         int    `json:"id"`
	RecordType string `json:"recordType"`
	Method     string `json:"method"`
}

// DocumentHandler : returns a document based on an id it is given
func DocumentHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, _ := strconv.Atoi(ps.ByName("id"))

	document := Document{
		id,
		"Document",
		"GET",
	}

	log.Println(document)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(document); err != nil {
		log.Fatal(err)
	}
}

package httplib

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func SendResponse(w http.ResponseWriter, statusCode int, body any) {
	data, err := json.Marshal(body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		msg := fmt.Sprintf("error sending request, fail marshal body: %v", err)
		_, err := w.Write([]byte(msg))
		if err != nil {
			log.Print(err) // TODO when the logger is added change to log.Errorf
		}
	}
	w.WriteHeader(statusCode)
	_, err = w.Write(data)
	if err != nil {
		log.Print(err) // TODO when the logger is added change to log.Errorf
	}
}

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/mskrha/gosmtp"
)

/*
	Handler function for the send API
*/
func handleSend(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	debug("handleSend: Client connected from %s", getClient(r))

	applyCors(w)

	if r.URL.RequestURI() != "/send/" {
		w.Header().Set("Location", "/send/")
		w.WriteHeader(http.StatusMovedPermanently)
		return
	}

	if r.Method != "POST" {
		debug("handleSend: Method %s not allowed!", r.Method)
		allowed(w, "POST")
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var msg gosmtp.Message
	err = json.Unmarshal(body, &msg)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := config.SMTP.server.Send(msg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	debug("handleSend: Message queued on the SMTP proxy with ID: %s", id)
	http.Error(w, fmt.Sprintf("Message queued on the SMTP proxy with ID: %s", id), http.StatusOK)

	return
}

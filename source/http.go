package main

import (
	"net/http"
	"strings"
)

/*
	Prepare and start the HTTP server
*/
func startHttp() {
	/*
		Response with '400 Bad Request' to all unhandled URLs
	*/
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		return
	})

	/*
		Used for health checking
	*/
	http.HandleFunc("/ping/", func(w http.ResponseWriter, r *http.Request) {
		return
	})

	http.HandleFunc("/send/", handleSend)

	/*
		Start the HTTP server
	*/
	debug("HTTP listening on %s", config.HTTP.listen)
	if err := http.ListenAndServe(config.HTTP.listen, nil); err != nil {
		panic(err)
	}
}

/*
	Return the HTTP client address

	X-Forwarded-For if set, else TCP remote address
*/
func getClient(r *http.Request) string {
	if len(r.Header.Get("X-Forwarded-For")) == 0 {
		return strings.Split(r.RemoteAddr, ":")[0]
	} else {
		return r.Header.Get("X-Forwarded-For")
	}
}

/*
	Apply the Cross-origin resource sharing header
*/
func applyCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
}

func allowed(w http.ResponseWriter, m string) {
	w.Header().Set("Allow", m)
	w.WriteHeader(http.StatusMethodNotAllowed)
}

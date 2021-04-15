package main

import (
	"github.com/mskrha/gosmtp"
)

/*
	Global configuration
*/
type configGlobal struct {
	Debug bool       `json:"debug"`
	HTTP  configHTTP `json:"http"`
	SMTP  configSmtp `json:"smtp"`

	/*
		Local hostname
	*/
	host string
}

/*
	HTTP server configuration
*/
type configHTTP struct {
	IP   string `json:"ip"`
	Port uint   `json:"port`

	listen string
}

/*
	SMTP relay configuration
*/
type configSmtp struct {
	Host  string `json:"host"`
	Port  uint   `json:"uint"`
	Agent string `json:"agent"`

	server *gosmtp.SMTP
}

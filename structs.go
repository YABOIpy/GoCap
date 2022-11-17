package gocap

import (
	"net/http"
	"crypto/tls"
)


type capresp struct {
	Captcha     interface{}
	solution    interface{}
	TaskID		interface{}
}

type structs struct {
	UserAgent	string
	WebKey		string
	WebURL		string
}


var (
	Client = http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				MaxVersion: tls.VersionTLS13,
			},
		},
	}
)
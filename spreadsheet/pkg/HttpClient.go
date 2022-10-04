package pkg

import (
	"io/ioutil"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func HttpClient(credentialFilePath string) (*http.Client, error) {
	data, err := ioutil.ReadFile(credentialFilePath)
	if err != nil {
		return nil, err
	}
	conf, err := google.JWTConfigFromJSON(data, "https://www.googleapis.com/auth/spreadsheets")
	if err != nil {
		return nil, err
	}

	return conf.Client(oauth2.NoContext), nil
}

/*
Copyright 2018 Comcast Cable Communications Management, LLC
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package vinyldns

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
)

type testToolsConfig struct {
	endpoint string
	code     int
	body     string
}

func readFile(file string) (string, error) {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func testTools(configs []testToolsConfig) (*httptest.Server, *Client) {
	host := "http://host.com"

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		reqURL, _ := url.Parse(r.RequestURI)
		reqPath := reqURL.Path
		reqQuery := reqURL.Query()

		for _, c := range configs {

			expURL, _ := url.Parse(c.endpoint)
			expPath := expURL.Path
			expQuery := expURL.Query()

			if reqPath != expPath {
				continue
			}

			if reflect.DeepEqual(reqQuery, expQuery) {
				w.WriteHeader(c.code)
				w.Header().Set("Content-Type", "application/json")
				fmt.Fprintf(w, c.body)
				return
			}
		}

		fmt.Printf("Requested: %s\n", r.RequestURI)
		http.Error(w, "not found", http.StatusNotFound)
		return
	}))

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		Proxy: func(req *http.Request) (*url.URL, error) {
			return url.Parse(server.URL)
		},
	}

	client := &Client{
		"accessToken",
		"secretToken",
		host,
		&http.Client{Transport: tr},
		"go-vinyldns testing",
	}

	return server, client
}

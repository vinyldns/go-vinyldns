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
	"net/http"
	"os"
)

// ClientConfiguration represents the vinyldns client configuration.
type ClientConfiguration struct {
	AccessKey string
	SecretKey string
	Host      string
}

// NewConfigFromEnv creates a new ClientConfiguration
// using environment variables.
func NewConfigFromEnv() ClientConfiguration {
	return ClientConfiguration{
		os.Getenv("VINYLDNS_ACCESS_KEY"),
		os.Getenv("VINYLDNS_SECRET_KEY"),
		os.Getenv("VINYLDNS_HOST"),
	}
}

// Client is a vinyldns API client.
type Client struct {
	AccessKey  string
	SecretKey  string
	Host       string
	HTTPClient *http.Client
}

// NewClientFromEnv returns a Client configured via
// environment variables.
func NewClientFromEnv() *Client {
	return NewClient(NewConfigFromEnv())
}

// NewClient returns a new vinyldns client using
// the client ClientConfiguration it's passed.
func NewClient(config ClientConfiguration) *Client {
	return &Client{
		config.AccessKey,
		config.SecretKey,
		config.Host,
		&http.Client{},
	}
}

func logRequests() bool {
	return os.Getenv("VINYLDNS_LOG") != ""
}

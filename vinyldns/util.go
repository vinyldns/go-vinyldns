/*
Copyright 2026 Comcast Cable Communications Management, LLC
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
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	awsauth "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	awscred "github.com/aws/aws-sdk-go-v2/credentials"
)

func concat(arr []string) string {
	var buff bytes.Buffer

	for _, elem := range arr {
		buff.WriteString(elem)
	}

	return buff.String()
}

func concatStrs(delim string, str ...string) string {
	return strings.Join(str, delim)
}

func resourceRequest(c *Client, url, method string, body []byte, responseStruct interface{}) error {
	if logRequests() {
		fmt.Printf("Request url: \n\t%s\nrequest body: \n\t%s \n\n", url, string(body))
	}
	req, err := http.NewRequest(method, url, bytes.NewReader(body))
	if err != nil {
		return err
	}

	req.Header.Set("User-Agent", c.UserAgent)
	req.Header.Set("Content-Type", "application/json")

	signer := awsauth.NewSigner()
	creds := awscred.NewStaticCredentialsProvider(c.AccessKey, c.SecretKey, "")

	h := sha256.New()
	_, _ = io.Copy(h, bytes.NewReader(body))
	payloadHash := hex.EncodeToString(h.Sum(nil))
	err = signer.SignHTTP(nil, creds.Value, req, payloadHash, "VinylDNS", "us-east-1", time.Now())
	if err != nil {
		return err
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	bodyContents, err := io.ReadAll(resp.Body)
	if logRequests() {
		fmt.Printf("Response status: \n\t%d\nresponse body: \n\t%s \n\n", resp.StatusCode, bodyContents)
	}
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusAccepted {
		dError := &Error{}
		dError.RequestURL = url
		dError.RequestMethod = method
		dError.RequestBody = string(body)
		dError.ResponseCode = resp.StatusCode
		dError.ResponseBody = string(bodyContents)
		return dError
	}
	err = json.Unmarshal(bodyContents, responseStruct)
	if err != nil {
		return err
	}
	return nil
}

func resourceRequestRaw(c *Client, url, method string, body []byte) (int, string, error) {
	if logRequests() {
		fmt.Printf("Request url: \n\t%s\nrequest body: \n\t%s \n\n", url, string(body))
	}
	req, err := http.NewRequest(method, url, bytes.NewReader(body))
	if err != nil {
		return 0, "", err
	}

	req.Header.Set("User-Agent", c.UserAgent)
	req.Header.Set("Content-Type", "application/json")

	signer := awsauth.NewSigner()
	creds := awscred.NewStaticCredentialsProvider(c.AccessKey, c.SecretKey, "")

	h := sha256.New()
	_, _ = io.Copy(h, bytes.NewReader(body))
	payloadHash := hex.EncodeToString(h.Sum(nil))
	err = signer.SignHTTP(nil, creds.Value, req, payloadHash, "VinylDNS", "us-east-1", time.Now())
	if err != nil {
		return 0, "", err
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return 0, "", err
	}

	bodyContents, err := io.ReadAll(resp.Body)
	if logRequests() {
		fmt.Printf("Response status: \n\t%d\nresponse body: \n\t%s \n\n", resp.StatusCode, bodyContents)
	}
	if err != nil {
		return resp.StatusCode, "", err
	}
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusAccepted {
		dError := &Error{}
		dError.RequestURL = url
		dError.RequestMethod = method
		dError.RequestBody = string(body)
		dError.ResponseCode = resp.StatusCode
		dError.ResponseBody = string(bodyContents)
		return resp.StatusCode, "", dError
	}

	return resp.StatusCode, string(bodyContents), nil
}

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
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestRecordSetOwnershipTransferRequest(t *testing.T) {
	body := runOwnershipTransferTest(t, OwnershipTransferStatusRequested, false)

	if !bytes.Contains(body, []byte(`"ownershipTransferStatus":"Requested"`)) {
		t.Error("expected ownershipTransferStatus to be Requested")
	}
	if !bytes.Contains(body, []byte(`"requestedOwnerGroupId":"requested-group-id"`)) {
		t.Error("expected requestedOwnerGroupId in request body")
	}
	if !bytes.Contains(body, []byte(`"ownerGroupId":"owner-group-id"`)) {
		t.Error("expected ownerGroupId to remain unchanged for request")
	}
}

func TestRecordSetOwnershipTransferApprove(t *testing.T) {
	body := runOwnershipTransferTest(t, OwnershipTransferStatusManuallyApproved, true)

	if !bytes.Contains(body, []byte(`"ownershipTransferStatus":"ManuallyApproved"`)) {
		t.Error("expected ownershipTransferStatus to be ManuallyApproved")
	}
	if !bytes.Contains(body, []byte(`"requestedOwnerGroupId":"requested-group-id"`)) {
		t.Error("expected requestedOwnerGroupId in request body")
	}
	if !bytes.Contains(body, []byte(`"ownerGroupId":"requested-group-id"`)) {
		t.Error("expected ownerGroupId to be updated for approve")
	}
}

func TestRecordSetOwnershipTransferReject(t *testing.T) {
	body := runOwnershipTransferTest(t, OwnershipTransferStatusManuallyRejected, false)

	if !bytes.Contains(body, []byte(`"ownershipTransferStatus":"ManuallyRejected"`)) {
		t.Error("expected ownershipTransferStatus to be ManuallyRejected")
	}
	if !bytes.Contains(body, []byte(`"requestedOwnerGroupId":"requested-group-id"`)) {
		t.Error("expected requestedOwnerGroupId in request body")
	}
}

func TestRecordSetOwnershipTransferCancel(t *testing.T) {
	body := runOwnershipTransferTest(t, OwnershipTransferStatusCancelled, false)

	if !bytes.Contains(body, []byte(`"ownershipTransferStatus":"Cancelled"`)) {
		t.Error("expected ownershipTransferStatus to be Cancelled")
	}
	if !bytes.Contains(body, []byte(`"requestedOwnerGroupId":"requested-group-id"`)) {
		t.Error("expected requestedOwnerGroupId in request body")
	}
}

func runOwnershipTransferTest(t *testing.T, status OwnershipTransferStatus, updateOwnerGroup bool) []byte {
	t.Helper()

	recordSetUpdateResponseJSON, err := readFile("test-fixtures/recordsets/recordset-update.json")
	if err != nil {
		t.Fatal(err)
	}

	var body []byte
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.RequestURI != "/zones/123/recordsets/456" {
			t.Fatalf("unexpected endpoint %s", r.RequestURI)
		}

		body, _ = io.ReadAll(r.Body)
		w.WriteHeader(http.StatusAccepted)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, recordSetUpdateResponseJSON)
	}))
	defer server.Close()

	client := newOwnershipTransferClient(server.URL)

	rs := &RecordSet{
		ZoneID:       "123",
		ID:           "456",
		Name:         "name",
		Type:         "CNAME",
		TTL:          200,
		OwnerGroupID: "owner-group-id",
		Records: []Record{{
			CName: "cname",
		}},
	}

	switch status {
	case OwnershipTransferStatusRequested:
		_, err = client.RecordSetOwnershipTransferRequest(rs, "requested-group-id")
	case OwnershipTransferStatusManuallyApproved:
		_, err = client.RecordSetOwnershipTransferApprove(rs, "requested-group-id")
	case OwnershipTransferStatusManuallyRejected:
		_, err = client.RecordSetOwnershipTransferReject(rs, "requested-group-id")
	case OwnershipTransferStatusCancelled:
		_, err = client.RecordSetOwnershipTransferCancel(rs, "requested-group-id")
	default:
		t.Fatalf("unexpected ownership transfer status %s", status)
	}

	if err != nil {
		t.Fatal(err)
	}

	if updateOwnerGroup && !bytes.Contains(body, []byte(`"ownerGroupId":"requested-group-id"`)) {
		t.Error("expected ownerGroupId to be updated for approval")
	}

	return body
}

func newOwnershipTransferClient(serverURL string) *Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		Proxy: func(req *http.Request) (*url.URL, error) {
			return url.Parse(serverURL)
		},
	}

	return &Client{
		AccessKey:  "accessToken",
		SecretKey:  "secretToken",
		Host:       "http://host.com",
		HTTPClient: &http.Client{Transport: tr},
		UserAgent:  "go-vinyldns testing",
	}
}

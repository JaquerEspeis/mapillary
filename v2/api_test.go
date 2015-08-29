// -*- Mode: Go; indent-tabs-mode: t -*-

/*
 * Copyright (C) 2015 JaquerEspeis
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License version 3 as
 * published by the Free Software Foundation.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 */

package api

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	check "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { check.TestingT(t) }

type APIv2Suite struct{}

var _ = check.Suite(&APIv2Suite{})

func (s *APIv2Suite) TestNewClientMustUseDefaultURL(c *check.C) {
	client := NewClient("dummy-id")
	c.Assert(client.URL.String(), check.Equals, "https://a.mapillary.com/v2",
		check.Commentf("Wrong URL"))
}

func (s *APIv2Suite) TestNewClientMustUseID(c *check.C) {
	client := NewClient("test-id")
	c.Assert(client.ID, check.Equals, "test-id", check.Commentf("Wrong ID"))
}

func (s *APIv2Suite) TestRequestMustAppendPathToURL(c *check.C) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c.Assert(r.URL.Path, check.Equals, "/test-version/test-path", check.Commentf("Wrong path"))
		io.WriteString(w, "{}")
	}))
	c.Assert(mockServer, check.NotNil, check.Commentf("Error creating mock server"))
	defer mockServer.Close()

	mockServerURL, err := url.Parse(mockServer.URL)
	c.Assert(err, check.IsNil, check.Commentf("Error parsing URL: %s", err))

	client := NewClient("dummy-id")
	client.URL = *mockServerURL
	client.URL.Path = "test-version"

	var r interface{}
	err = client.Request("dummy-method", "test-path", url.Values{}, &r)
	c.Assert(err, check.IsNil, check.Commentf("Error on request: %s", err))
}

func (s *APIv2Suite) TestRequestMustAppendClientIDToURL(c *check.C) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c.Assert(r.URL.Query()["client_id"], check.DeepEquals, []string{"test-id"},
			check.Commentf("Wrong client ID query"))

		io.WriteString(w, "{}")
	}))
	c.Assert(mockServer, check.NotNil, check.Commentf("Error creating mock server"))
	defer mockServer.Close()

	mockServerURL, err := url.Parse(mockServer.URL)
	c.Assert(err, check.IsNil, check.Commentf("Error parsing URL: %s", err))

	client := NewClient("test-id")
	client.URL = *mockServerURL

	var r interface{}
	err = client.Request("dummy-method", "test-path", url.Values{}, &r)
	c.Assert(err, check.IsNil, check.Commentf("Error on request: %s", err))
}

func (s *APIv2Suite) TestRequestMustAppendParamsToURL(c *check.C) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c.Check(r.URL.Query()["test-param1"], check.DeepEquals, []string{"test-value1"},
			check.Commentf("Wrong param value"))
		c.Check(r.URL.Query()["test-param2"], check.DeepEquals, []string{"test-value2"},
			check.Commentf("Wrong param value"))

		io.WriteString(w, "{}")
	}))
	c.Assert(mockServer, check.NotNil, check.Commentf("Error creating mock server"))
	defer mockServer.Close()

	mockServerURL, err := url.Parse(mockServer.URL)
	c.Assert(err, check.IsNil, check.Commentf("Error parsing URL: %s", err))

	client := NewClient("test-id")
	client.URL = *mockServerURL

	var r interface{}
	params := url.Values{
		"test-param1": {"test-value1"},
		"test-param2": {"test-value2"},
	}
	err = client.Request("dummy-method", "test-path", params, &r)

	c.Assert(err, check.IsNil, check.Commentf("Error on request: %s", err))
}

func (s *APIv2Suite) TestRequestMustReturnJSON(c *check.C) {
	testJSON := map[string]interface{}{
		"test-param1": "test-value1",
		"test-param2": "test-value2",
	}

	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b, err := json.Marshal(testJSON)
		c.Assert(err, check.IsNil, check.Commentf("Error encoding the JSON: %s", err))
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
	}))
	c.Assert(mockServer, check.NotNil, check.Commentf("Error creating mock server"))
	defer mockServer.Close()

	mockServerURL, err := url.Parse(mockServer.URL)
	c.Assert(err, check.IsNil, check.Commentf("Error parsing URL: %s", err))

	client := NewClient("test-id")
	client.URL = *mockServerURL

	var r interface{}
	params := url.Values{
		"test-param1": {"test-value1"},
		"test-param2": {"test-value2"},
	}
	err = client.Request("dummy-method", "test-path", params, &r)
	c.Assert(err, check.IsNil, check.Commentf("Error on request: %s", err))

	c.Assert(r, check.DeepEquals, testJSON)
}

func (s *APIv2Suite) TestGetMustSetRequestMethod(c *check.C) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c.Assert(r.Method, check.Equals, "GET", check.Commentf("Wrong method"))
		io.WriteString(w, "{}")
	}))
	c.Assert(mockServer, check.NotNil, check.Commentf("Error creating mock server"))
	defer mockServer.Close()

	mockServerURL, err := url.Parse(mockServer.URL)
	c.Assert(err, check.IsNil, check.Commentf("Error parsing URL: %s", err))

	client := NewClient("dummy-id")
	client.URL = *mockServerURL

	var r interface{}
	err = client.Get("test-path", url.Values{}, &r)
	c.Assert(err, check.IsNil, check.Commentf("Error on get: %s", err))
}

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

// Package api implements a client for Mapillary API version 2.
package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	pkgpath "path"
)

var mapillaryV2URL = url.URL{
	Scheme: "https",
	Host:   "a.mapillary.com",
	Path:   "v2",
}

// Client is a client to make requests to the Mapillary API version 2.
type Client struct {
	// HTTPClient is the client to perform the HTTP requests.
	HTTPClient *http.Client
	// URL is the location of the Mapillary API.
	URL url.URL
	// ID is the identifier of the client.
	ID string
}

// NewClient constructs a new API client with the default URL.
func NewClient(ID string) *Client {
	return &Client{HTTPClient: &http.Client{}, URL: mapillaryV2URL, ID: ID}
}

// Request performs an HTTP Request to the Mapillary API.
func (c *Client) Request(method, path string, params url.Values, response interface{}) error {
	params.Add("client_id", c.ID)
	c.URL.Path = pkgpath.Join(c.URL.Path, path)
	c.URL.RawQuery = params.Encode()
	request, err := http.NewRequest(method, c.URL.String(), nil)
	if err != nil {
		return err
	}
	r, err := c.HTTPClient.Do(request)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, &response)
}

// Get performs an HTTP Request to the Mapillary API.
func (c *Client) Get(path string, params url.Values, response interface{}) error {
	return c.Request("GET", path, params, response)
}

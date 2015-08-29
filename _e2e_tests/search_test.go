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

package tests

import (
	"encoding/json"
	"image"
	// Imported to allow image.Decode to understand JPEG formatted images.
	_ "image/jpeg"
	"io/ioutil"
	"net/url"
	"os"
	"path"
	"testing"

	check "gopkg.in/check.v1"

	"github.com/JaquerEspeis/mapillary"
	"github.com/JaquerEspeis/mapillary/v2"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { check.TestingT(t) }

type SearchSuite struct{}

var _ = check.Suite(&SearchSuite{})

type conf struct {
	ClientID string
}

func clientID(c *check.C) string {
	confFile, err := os.Open("conf.json")
	c.Assert(err, check.IsNil, check.Commentf("Error opening conf file: %s", err))
	defer confFile.Close()
	var config conf
	confFileContents, err := ioutil.ReadAll(confFile)
	c.Assert(err, check.IsNil, check.Commentf("Error reading the conf file: %s", err))
	err = json.Unmarshal(confFileContents, &config)
	c.Assert(err, check.IsNil, check.Commentf("Error unmarshaling JSON config: %s", err))
	return config.ClientID
}

func (s *SearchSuite) TestSearchImRandomSelected(c *check.C) {
	client := api.NewClient(clientID(c))
	var response api.GetSearchImRandomSelect
	err := client.Get("search/im/randomselected", url.Values{}, &response)
	c.Assert(err, check.IsNil, check.Commentf("Error on request: %s", err))

	// Check that the key corresponds to an actual image and can be downloded.
	tmp, err := ioutil.TempDir("", "")
	c.Assert(err, check.IsNil, check.Commentf("Error creataing temp dir: %s", err))
	imgPath := path.Join(tmp, "image.jpeg")
	err = mapillary.DownloadImage(response.Key, 320, imgPath)
	c.Assert(err, check.IsNil, check.Commentf("Error downloading image: %s", err))

	imgFile, err := os.Open(imgPath)
	c.Assert(err, check.IsNil, check.Commentf("Error opening image: %s", err))
	defer imgFile.Close()
	_, f, err := image.Decode(imgFile)
	c.Assert(err, check.IsNil, check.Commentf("Error decoding image: %s", err))
	c.Assert(f, check.Equals, "jpeg")
}

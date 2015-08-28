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

package mapillary

import (
	"testing"

	check "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { check.TestingT(t) }

type MapillarySuite struct{}

var _ = check.Suite(&MapillarySuite{})

func (s *MapillarySuite) TestGetImageURL(c *check.C) {
	url, err := GetImageURL("test-image", 10)
	c.Assert(err, check.IsNil, check.Commentf("Error getting image: %s", err))
	c.Assert(url, check.Equals, "https://d1cuyjsrcm0gby.cloudfront.net/test-image/thumb-10.jpg")
}

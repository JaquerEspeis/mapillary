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

// Package mapillary implements functions to get images from the Mapillary storage.
package mapillary

import "fmt"

// GetImageURL returns the URL of an image in the Mapillary storage.
func GetImageURL(key string, size int) (string, error) {
	return fmt.Sprintf("https://d1cuyjsrcm0gby.cloudfront.net/%s/thumb-%d.jpg", key, size), nil
}

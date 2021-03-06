// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package yandex provides constants for using OAuth2 to access Yandex APIs.
package icheck // import "golang.org/x/oauth2/icheck"

import (
	"golang.org/x/oauth2"
)

// Endpoint is the Yandex OAuth 2.0 endpoint.
var Endpoint = oauth2.Endpoint{
	AuthURL:  "http://sandbox.icheck.com.vn:4347/authorize",
	TokenURL: "http://sandbox.icheck.com.vn:4347/token",
}

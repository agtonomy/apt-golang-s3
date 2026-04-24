// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package method

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
)

var errEmptyRegion = errors.New("region is required")

// s3EndpointURL returns the default S3 endpoint URL for the given region.
// us-east-1 uses the legacy global hostname; cn-* regions use the .com.cn
// suffix; everything else follows s3.<region>.amazonaws.com.
func s3EndpointURL(region string) (*url.URL, error) {
	if region == "" {
		return nil, errEmptyRegion
	}

	var host string
	switch {
	case region == "us-east-1":
		host = "s3.amazonaws.com"
	case strings.HasPrefix(region, "cn-"):
		host = fmt.Sprintf("s3.%s.amazonaws.com.cn", region)
	default:
		host = fmt.Sprintf("s3.%s.amazonaws.com", region)
	}

	return &url.URL{Scheme: "https", Host: host}, nil
}

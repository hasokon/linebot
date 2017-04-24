// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

// Command label uses the Vision API's label detection capabilities to find a label based on an image's content.
package main

import (
	"io"
	"os"

	// [START imports]
	"cloud.google.com/go/vision"
	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	// [END imports]
	"errors"
)

// FindLabels gets labels from the Vision API for an image at the given file path.
func FindLabels(c io.ReadCloser) ([]string, error) {
	// [START init]
	json := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	ctx := context.Background()
	jwtConfig, err := google.JWTConfigFromJSON([]byte(json), vision.Scope)
	if err != nil {
		return nil, errors.New("google.JWTConfigFromJSON :" + err.Error() + "\n" + json)
	}

	ts := jwtConfig.TokenSource(ctx)

	// Create the client.
	client, err := vision.NewClient(ctx, option.WithTokenSource(ts))
	if err != nil {
		return nil, err
	}
	// [END init]

	// [START request]
	// Perform the request
	image, err := vision.NewImageFromReader(c)
	if err != nil {
		return nil, err
	}

	annotations, err := client.DetectLabels(ctx, image, 10)
	if err != nil {
		return nil, err
	}
	// [END request]
	// [START transform]
	var labels []string
	for _, annotation := range annotations {
		labels = append(labels, annotation.Description)
	}
	return labels, nil
	// [END transform]
}

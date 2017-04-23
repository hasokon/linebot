// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

// Command label uses the Vision API's label detection capabilities to find a label based on an image's content.
package main

import (
	"bytes"
	"io/ioutil"

	// [START imports]
	"cloud.google.com/go/vision"
	"golang.org/x/net/context"
	// [END imports]
)

// findLabels gets labels from the Vision API for an image at the given file path.
func FindLabels(b []byte) ([]string, error) {
	// [START init]
	ctx := context.Background()

	// Create the client.
	client, err := vision.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	// [END init]

	// [START request]
	// Perform the request
	reader, err := bytes.NewReader(b)
	if err != nil {
		return nil, err
	}

	readcloser := ioutil.NopCloser(reader)

	image, err := NewImageFromReader(readcloser)
	if err != nil {
		return nil, err
	}

	annotations, err := client.DetectLabels(ctx, usableImage, 10)
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

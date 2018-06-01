// Copyright 2017 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

// Command detect uses the Vision API's label detection capabilities to find a label based on an image's content.
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {

	// samples := []struct {
	// 	name       string
	// 	local,
	// 	uri func(io.Writer, string) error
	// }{
	// 	// {"detectFaces", detectFaces, detectFacesURI},
	// 	// {"detectLabels", detectLabels, detectLabelsURI},
	// 	// {"detectLandmarks", detectLandmarks, detectLandmarksURI},
	// 	// {"detectText", detectText, detectTextURI},
	// 	// {"detectDocumentText", detectDocumentText, detectDocumentTextURI},
	// 	{"detectDocumentText", detectDocumentTextCustom, detectDocumentTextURI},
	// 	// {"detectLogos", detectLogos, detectLogosURI},
	// 	// {"detectProperties", detectProperties, detectPropertiesURI},
	// 	// {"detectCropHints", detectCropHints, detectCropHintsURI},
	// 	// {"detectWeb", detectWeb, detectWebURI},
	// 	// {"detectWebGeo", detectWebGeo, detectWebGeoURI},
	// 	// {"detectSafeSearch", detectSafeSearch, detectSafeSearchURI},
	// }

	samples := []struct {
		name string
		uri  func(io.Writer, string) error
	}{
		// {"detectFaces", detectFaces, detectFacesURI},
		// {"detectLabels", detectLabels, detectLabelsURI},
		// {"detectLandmarks", detectLandmarks, detectLandmarksURI},
		// {"detectText", detectText, detectTextURI},
		// {"detectDocumentText", detectDocumentText, detectDocumentTextURI},
		{"detectDocumentText", detectDocumentTextCustom},
		// {"detectLogos", detectLogos, detectLogosURI},
		// {"detectProperties", detectProperties, detectPropertiesURI},
		// {"detectCropHints", detectCropHints, detectCropHintsURI},
		// {"detectWeb", detectWeb, detectWebURI},
		// {"detectWebGeo", detectWebGeo, detectWebGeoURI},
		// {"detectSafeSearch", detectSafeSearch, detectSafeSearchURI},
	}

	files, err := ioutil.ReadDir("images")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".jpg") {
			path := "images/" + file.Name()
			for _, sample := range samples {
				fmt.Println("---", path)
				var err error
				err = sample.uri(os.Stdout, path)
				if err != nil {
					fmt.Println("Error:", err)
				}
			}
		}

	}

}

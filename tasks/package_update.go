// Copyright 2025 The Outline Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"log"
	"os"
	"text/template"
)

type packageData struct {
	Tag      string
	Checksum string
}

func main() {
	if len(os.Args) != 3 {
		log.Fatal("Usage: go run package_update.go <tag> <checksum>")
	}
	tag := os.Args[1]
	checksum := os.Args[2]

	data := packageData{
		Tag:      tag,
		Checksum: checksum,
	}

	tmpl, err := template.ParseFiles("Package.swift.template")
	if err != nil {
		log.Fatalf("Failed to parse template file: %v", err)
	}

	outputFile, err := os.Create("Package.swift")
	if err != nil {
		log.Fatalf("Failed to create Package.swift: %v", err)
	}
	defer outputFile.Close()

	if err := tmpl.Execute(outputFile, data); err != nil {
		log.Fatalf("Failed to execute template: %v", err)
	}

	log.Println("Package.swift generated successfully.")
}

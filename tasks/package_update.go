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

	packageTmpl, err := template.ParseFiles("Package.swift.template")
	if err != nil {
		log.Fatalf("Failed to parse template file: %v", err)
	}

	packageOutputFile, err := os.Create("Package.swift")
	if err != nil {
		log.Fatalf("Failed to create Package.swift: %v", err)
	}
	defer packageOutputFile.Close()

	if err := packageTmpl.Execute(packageOutputFile, data); err != nil {
		log.Fatalf("Failed to execute template: %v", err)
	}

	gradleTmpl, err := template.ParseFiles("build.gradle.kts.template")
	if err != nil {
		log.Fatalf("Failed to parse template file: %v", err)
	}

	gradleOutputFile, err := os.Create("build.gradle.kts")
	if err != nil {
		log.Fatalf("Failed to create build.gradle.kts: %v", err)
	}
	defer gradleOutputFile.Close()

	if err := gradleTmpl.Execute(gradleOutputFile, data); err != nil {
		log.Fatalf("Failed to execute template: %v", err)
	}

	log.Println("Package.swift and build.gradle.kts generated successfully.")
}

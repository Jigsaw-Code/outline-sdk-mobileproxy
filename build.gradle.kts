// Copyright 2025 The Outline Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

plugins {
  id "maven-publish"
}

val aarFile = layout.buildDirectory.file("mobileproxy.aar")

tasks.register<Exec>("downloadRelease") {
  commandLine(
    "curl",
    "-L",
    "-o", aarFile.get().asFile.path,
    "https://github.com/Jigsaw-Code/outline-sdk-mobileproxy/releases/download/${project.version}/mobileproxy.aar"
  )
}

publishing {
  publications {
    create<MavenPublication>("release") {
      artifact(aarFile) {
        builtBy(tasks.named("downloadRelease"))
      }
    }
  }
}

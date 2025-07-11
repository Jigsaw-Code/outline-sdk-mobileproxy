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

// !!! THIS IS A PLACEHOLDER, NOT A VALID PACKAGE FILE !!!

// swift-tools-version:5.3
import PackageDescription

let package = Package(
    name: "OutlineSDK-Mobileproxy",
    platforms: [
        .iOS(.v11)
    ],
    products: [
        .library(
            name: "OutlineSDK-Mobileproxy",
            targets: ["mobileproxy"])
    ],
    targets: [
        .binaryTarget(
            name: "mobileproxy",
            url: "https://github.com/Jigsaw-Code/outline-sdk-mobileproxy/releases/download/0.0.0/mobileproxy.xcframework.zip",
            checksum: "fc7e8d5e077690fc410d5379ec178176"
        )
    ]
)

# Outline SDK Mobileproxy

This repository provides pre-compiled binaries of the Outline SDK Mobileproxy for Android and iOS. See the [Outline SDK](https://github.com/Jigsaw-Code/outline-sdk) for the library's source code.

## Release Process

To create a new release, you must run a manual GitHub Action. Follow these steps:

1.  Navigate to the **Actions** tab.
2.  Select the **Draft Mobileproxy Release** workflow.
3.  Click on the **Run workflow** dropdown.
4.  Enter the **Outline SDK tag** you want to base the release on (e.g., `x/v*.*.*`).
5.  Click the **Run workflow** button. This will trigger a new workflow run that builds the mobileproxy artifacts, updates the `Package.swift` and `build.gradle.kts` files, and creates a new draft release with the generated binaries.
6.  Update the release notes and verify that the library works manually.
7.  Publish the release!

## Integration

## Android

We use JitPack to distribute the Android library. To integrate it into your app, follow these steps:

1.  Add the JitPack repository to your root `build.gradle` file:
    ```groovy
    allprojects {
        repositories {
            ...
            maven { url 'https://jitpack.io' }
        }
    }
    ```
2.  Add the dependency to your app's `build.gradle` file:
    ```groovy
    dependencies {
        implementation 'com.github.Jigsaw-Code:outline-sdk-mobileproxy:<version>'
    }
    ```

### iOS

You can add the Mobileproxy to your Xcode project using Swift Package Manager.

1.  In Xcode, go to **File \> Add Packages...**
2.  Enter the repository URL: `https://github.com/Jigsaw-Code/outline-sdk-mobileproxy`
3.  Select the desired version.

## License

This repository is licensed under the Apache License 2.0.

**Note:** The Psiphon extension is not supported as its code is under the GPL license.

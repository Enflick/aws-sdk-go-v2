# v1.4.4 (2024-05-23)

* No change notes available for this release.

# v1.4.3 (2024-05-16)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.4.2 (2024-05-15)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.4.1 (2024-05-08)

* **Bug Fix**: GoDoc improvement

# v1.4.0 (2024-04-03)

* **Feature**: The release includes a public SDK for AWS Clean Rooms ML APIs, making them globally available to developers worldwide.

# v1.3.4 (2024-03-29)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.3.3 (2024-03-18)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.3.2 (2024-03-07)

* **Bug Fix**: Remove dependency on go-cmp.
* **Dependency Update**: Updated to the latest SDK module versions

# v1.3.1 (2024-02-23)

* **Bug Fix**: Move all common, SDK-side middleware stack ops into the service client module to prevent cross-module compatibility issues in the future.
* **Dependency Update**: Updated to the latest SDK module versions

# v1.3.0 (2024-02-22)

* **Feature**: Add middleware stack snapshot tests.

# v1.2.2 (2024-02-21)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.2.1 (2024-02-20)

* **Bug Fix**: When sourcing values for a service's `EndpointParameters`, the lack of a configured region (i.e. `options.Region == ""`) will now translate to a `nil` value for `EndpointParameters.Region` instead of a pointer to the empty string `""`. This will result in a much more explicit error when calling an operation instead of an obscure hostname lookup failure.

# v1.2.0 (2024-02-13)

* **Feature**: Bump minimum Go version to 1.20 per our language support policy.
* **Dependency Update**: Updated to the latest SDK module versions

# v1.1.4 (2024-01-04)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.1.3 (2023-12-08)

* **Bug Fix**: Reinstate presence of default Retryer in functional options, but still respect max attempts set therein.

# v1.1.2 (2023-12-07)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.1.1 (2023-12-06)

* **Bug Fix**: Restore pre-refactor auth behavior where all operations could technically be performed anonymously.

# v1.1.0 (2023-12-05)

* **Feature**: Updated service title from cleanroomsml to CleanRoomsML.

# v1.0.2 (2023-12-01)

* **Bug Fix**: Correct wrapping of errors in authentication workflow.
* **Bug Fix**: Correctly recognize cache-wrapped instances of AnonymousCredentials at client construction.
* **Dependency Update**: Updated to the latest SDK module versions

# v1.0.1 (2023-11-30)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.0.0 (2023-11-29)

* **Release**: New AWS service client module
* **Feature**: Public Preview SDK release of AWS Clean Rooms ML APIs
* **Dependency Update**: Updated to the latest SDK module versions


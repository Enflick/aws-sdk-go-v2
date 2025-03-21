# v1.1.8 (2024-05-23)

* No change notes available for this release.

# v1.1.7 (2024-05-16)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.1.6 (2024-05-15)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.1.5 (2024-05-08)

* **Bug Fix**: GoDoc improvement

# v1.1.4 (2024-03-29)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.1.3 (2024-03-18)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.1.2 (2024-03-07)

* **Bug Fix**: Remove dependency on go-cmp.
* **Dependency Update**: Updated to the latest SDK module versions

# v1.1.1 (2024-02-23)

* **Bug Fix**: Move all common, SDK-side middleware stack ops into the service client module to prevent cross-module compatibility issues in the future.
* **Dependency Update**: Updated to the latest SDK module versions

# v1.1.0 (2024-02-22)

* **Feature**: Add middleware stack snapshot tests.

# v1.0.2 (2024-02-21)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.0.1 (2024-02-20)

* **Bug Fix**: When sourcing values for a service's `EndpointParameters`, the lack of a configured region (i.e. `options.Region == ""`) will now translate to a `nil` value for `EndpointParameters.Region` instead of a pointer to the empty string `""`. This will result in a much more explicit error when calling an operation instead of an obscure hostname lookup failure.

# v1.0.0 (2024-02-15)

* **Release**: New AWS service client module
* **Feature**: This is the initial SDK release for AWS Artifact. AWS Artifact provides on-demand access to compliance and third-party compliance reports. This release includes access to List and Get reports, along with their metadata. This release also includes access to AWS Artifact notifications settings.


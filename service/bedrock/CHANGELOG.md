# v1.8.5 (2024-05-23)

* No change notes available for this release.

# v1.8.4 (2024-05-16)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.8.3 (2024-05-15)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.8.2 (2024-05-08)

* **Bug Fix**: GoDoc improvement

# v1.8.1 (2024-05-01)

* No change notes available for this release.

# v1.8.0 (2024-04-23)

* **Feature**: This release introduces Model Evaluation and Guardrails for Amazon Bedrock.

# v1.7.7 (2024-04-10)

* No change notes available for this release.

# v1.7.6 (2024-04-04)

* No change notes available for this release.

# v1.7.5 (2024-04-02)

* No change notes available for this release.

# v1.7.4 (2024-03-29)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.7.3 (2024-03-18)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.7.2 (2024-03-07)

* **Bug Fix**: Remove dependency on go-cmp.
* **Dependency Update**: Updated to the latest SDK module versions

# v1.7.1 (2024-02-23)

* **Bug Fix**: Move all common, SDK-side middleware stack ops into the service client module to prevent cross-module compatibility issues in the future.
* **Dependency Update**: Updated to the latest SDK module versions

# v1.7.0 (2024-02-22)

* **Feature**: Add middleware stack snapshot tests.

# v1.6.2 (2024-02-21)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.6.1 (2024-02-20)

* **Bug Fix**: When sourcing values for a service's `EndpointParameters`, the lack of a configured region (i.e. `options.Region == ""`) will now translate to a `nil` value for `EndpointParameters.Region` instead of a pointer to the empty string `""`. This will result in a much more explicit error when calling an operation instead of an obscure hostname lookup failure.

# v1.6.0 (2024-02-13)

* **Feature**: Bump minimum Go version to 1.20 per our language support policy.
* **Dependency Update**: Updated to the latest SDK module versions

# v1.5.7 (2024-01-04)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.5.6 (2023-12-22)

* No change notes available for this release.

# v1.5.5 (2023-12-08)

* **Bug Fix**: Reinstate presence of default Retryer in functional options, but still respect max attempts set therein.

# v1.5.4 (2023-12-07)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.5.3 (2023-12-06)

* **Bug Fix**: Restore pre-refactor auth behavior where all operations could technically be performed anonymously.

# v1.5.2 (2023-12-01)

* **Bug Fix**: Correct wrapping of errors in authentication workflow.
* **Bug Fix**: Correctly recognize cache-wrapped instances of AnonymousCredentials at client construction.
* **Dependency Update**: Updated to the latest SDK module versions

# v1.5.1 (2023-11-30)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.5.0 (2023-11-29)

* **Feature**: Expose Options() accessor on service clients.
* **Dependency Update**: Updated to the latest SDK module versions

# v1.4.0 (2023-11-28.2)

* **Feature**: This release adds support for customization types, model life cycle status and minor versions/aliases for model identifiers.
* **Dependency Update**: Updated to the latest SDK module versions

# v1.3.4 (2023-11-28)

* **Bug Fix**: Respect setting RetryMaxAttempts in functional options at client construction.

# v1.3.3 (2023-11-20)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.3.2 (2023-11-15)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.3.1 (2023-11-09)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.3.0 (2023-11-01)

* **Feature**: Adds support for configured endpoints via environment variables and the AWS shared configuration file.
* **Dependency Update**: Updated to the latest SDK module versions

# v1.2.0 (2023-10-31)

* **Feature**: **BREAKING CHANGE**: Bump minimum go version to 1.19 per the revised [go version support policy](https://aws.amazon.com/blogs/developer/aws-sdk-for-go-aligns-with-go-release-policy-on-supported-runtimes/).
* **Dependency Update**: Updated to the latest SDK module versions

# v1.1.5 (2023-10-23)

* No change notes available for this release.

# v1.1.4 (2023-10-12)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.1.3 (2023-10-06)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.1.2 (2023-10-05)

* No change notes available for this release.

# v1.1.1 (2023-10-04)

* No change notes available for this release.

# v1.1.0 (2023-10-02)

* **Feature**: Provisioned throughput feature with Amazon and third-party base models, and update validators for model identifier and taggable resource ARNs.

# v1.0.0 (2023-09-28)

* **Release**: New AWS service client module
* **Feature**: Model Invocation logging added to enable or disable logs in customer account. Model listing and description support added. Provisioned Throughput feature added. Custom model support added for creating custom models. Also includes list, and delete functions for custom model.


# v1.21.4 (2024-05-23)

* No change notes available for this release.

# v1.21.3 (2024-05-16)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.21.2 (2024-05-15)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.21.1 (2024-05-08)

* **Bug Fix**: GoDoc improvement

# v1.21.0 (2024-04-30)

* **Feature**: Add support for workflow sharing and dynamic run storage

# v1.20.0 (2024-04-11)

* **Feature**: This release adds support for retrieval of S3 direct access metadata on sequence stores and read sets, and adds support for SHA256up and SHA512up HealthOmics ETags.

# v1.19.4 (2024-03-29)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.19.3 (2024-03-18)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.19.2 (2024-03-07)

* **Bug Fix**: Remove dependency on go-cmp.
* **Dependency Update**: Updated to the latest SDK module versions

# v1.19.1 (2024-02-23)

* **Bug Fix**: Move all common, SDK-side middleware stack ops into the service client module to prevent cross-module compatibility issues in the future.
* **Dependency Update**: Updated to the latest SDK module versions

# v1.19.0 (2024-02-22)

* **Feature**: Add middleware stack snapshot tests.

# v1.18.2 (2024-02-21)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.18.1 (2024-02-20)

* **Bug Fix**: When sourcing values for a service's `EndpointParameters`, the lack of a configured region (i.e. `options.Region == ""`) will now translate to a `nil` value for `EndpointParameters.Region` instead of a pointer to the empty string `""`. This will result in a much more explicit error when calling an operation instead of an obscure hostname lookup failure.

# v1.18.0 (2024-02-16)

* **Feature**: Add new ClientOptions field to waiter config which allows you to extend the config for operation calls made by waiters.

# v1.17.0 (2024-02-13)

* **Feature**: Bump minimum Go version to 1.20 per our language support policy.
* **Dependency Update**: Updated to the latest SDK module versions

# v1.16.7 (2024-01-04)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.16.6 (2023-12-22)

* **Documentation**: Provides minor corrections and an updated description of APIs.

# v1.16.5 (2023-12-08)

* **Bug Fix**: Reinstate presence of default Retryer in functional options, but still respect max attempts set therein.

# v1.16.4 (2023-12-07)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.16.3 (2023-12-06)

* **Bug Fix**: Restore pre-refactor auth behavior where all operations could technically be performed anonymously.

# v1.16.2 (2023-12-01)

* **Bug Fix**: Correct wrapping of errors in authentication workflow.
* **Bug Fix**: Correctly recognize cache-wrapped instances of AnonymousCredentials at client construction.
* **Dependency Update**: Updated to the latest SDK module versions

# v1.16.1 (2023-11-30)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.16.0 (2023-11-29)

* **Feature**: Expose Options() accessor on service clients.
* **Dependency Update**: Updated to the latest SDK module versions

# v1.15.4 (2023-11-28.2)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.15.3 (2023-11-28)

* **Bug Fix**: Respect setting RetryMaxAttempts in functional options at client construction.

# v1.15.2 (2023-11-20)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.15.1 (2023-11-15)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.15.0 (2023-11-09.2)

* **Feature**: Support UBAM filetype for Omics Storage and make referenceArn optional

# v1.14.1 (2023-11-09)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.14.0 (2023-11-08)

* **Feature**: Adding Run UUID and Run Output URI: GetRun and StartRun API response has two new fields "uuid" and "runOutputUri".

# v1.13.0 (2023-11-01)

* **Feature**: Adds support for configured endpoints via environment variables and the AWS shared configuration file.
* **Dependency Update**: Updated to the latest SDK module versions

# v1.12.0 (2023-10-31)

* **Feature**: **BREAKING CHANGE**: Bump minimum go version to 1.19 per the revised [go version support policy](https://aws.amazon.com/blogs/developer/aws-sdk-for-go-aligns-with-go-release-policy-on-supported-runtimes/).
* **Dependency Update**: Updated to the latest SDK module versions

# v1.11.0 (2023-10-19)

* **Feature**: This change enables customers to retrieve failure reasons with detailed status messages for their failed runs

# v1.10.2 (2023-10-12)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.10.1 (2023-10-06)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.10.0 (2023-10-05)

* **Feature**: Add Etag Support for Omics Storage in ListReadSets and GetReadSetMetadata API

# v1.9.1 (2023-09-27)

* No change notes available for this release.

# v1.9.0 (2023-08-29)

* **Feature**: Add RetentionMode support for Runs.

# v1.8.3 (2023-08-21)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.8.2 (2023-08-18)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.8.1 (2023-08-17)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.8.0 (2023-08-14)

* **Feature**: This release provides support for annotation store versioning and cross account sharing for Omics Analytics

# v1.7.0 (2023-08-10)

* **Feature**: This release adds instanceType to GetRunTask & ListRunTasks responses.

# v1.6.2 (2023-08-07)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.6.1 (2023-08-01)

* No change notes available for this release.

# v1.6.0 (2023-07-31)

* **Feature**: Add CreationType filter for ListReadSets
* **Feature**: Adds support for smithy-modeled endpoint resolution. A new rules-based endpoint resolution will be added to the SDK which will supercede and deprecate existing endpoint resolution. Specifically, EndpointResolver will be deprecated while BaseEndpoint and EndpointResolverV2 will take its place. For more information, please see the Endpoints section in our Developer Guide.
* **Dependency Update**: Updated to the latest SDK module versions

# v1.5.3 (2023-07-28)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.5.2 (2023-07-26)

* **Documentation**: The service is renaming as a part of AWS Health.

# v1.5.1 (2023-07-13)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.5.0 (2023-06-28)

* **Feature**: Add Common Workflow Language (CWL) as a supported language for Omics workflows

# v1.4.2 (2023-06-15)

* No change notes available for this release.

# v1.4.1 (2023-06-13)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.4.0 (2023-05-11)

* **Feature**: This release provides support for Ready2Run and GPU workflows, an improved read set filter, the direct upload of read sets into Omics Storage, and annotation parsing for analytics stores.

# v1.3.3 (2023-05-04)

* No change notes available for this release.

# v1.3.2 (2023-04-24)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.3.1 (2023-04-13)

* No change notes available for this release.

# v1.3.0 (2023-04-10)

* **Feature**: Remove unexpected API changes.

# v1.2.3 (2023-04-07)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.2.2 (2023-03-21)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.2.1 (2023-03-10)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.2.0 (2023-02-28)

* **Feature**: Minor model changes to accomodate batch imports feature

# v1.1.4 (2023-02-22)

* **Bug Fix**: Prevent nil pointer dereference when retrieving error codes.

# v1.1.3 (2023-02-20)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.1.2 (2023-02-15)

* **Announcement**: When receiving an error response in restJson-based services, an incorrect error type may have been returned based on the content of the response. This has been fixed via PR #2012 tracked in issue #1910.
* **Bug Fix**: Correct error type parsing for restJson services.

# v1.1.1 (2023-02-03)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.1.0 (2023-01-05)

* **Feature**: Add `ErrorCodeOverride` field to all error structs (aws/smithy-go#401).

# v1.0.2 (2022-12-15)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.0.1 (2022-12-02)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.0.0 (2022-11-29.2)

* **Release**: New AWS service client module
* **Feature**: Amazon Omics is a new, purpose-built service that can be used by healthcare and life science organizations to store, query, and analyze omics data. The insights from that data can be used to accelerate scientific discoveries and improve healthcare.


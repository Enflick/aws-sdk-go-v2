# v1.34.1 (2024-05-23)

* No change notes available for this release.

# v1.34.0 (2024-05-22)

* **Feature**: This release adds support for enabling or disabling a data source configured as part of Zero-ETL integration with Amazon S3, by setting its status.

# v1.33.3 (2024-05-16)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.33.2 (2024-05-15)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.33.1 (2024-05-08)

* **Bug Fix**: GoDoc improvement

# v1.33.0 (2024-04-30)

* **Feature**: This release enables customers to create Route53 A and AAAA alias record types to point custom endpoint domain to OpenSearch domain's dualstack search endpoint.

# v1.32.4 (2024-03-29)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.32.3 (2024-03-18)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.32.2 (2024-03-07)

* **Bug Fix**: Remove dependency on go-cmp.
* **Dependency Update**: Updated to the latest SDK module versions

# v1.32.1 (2024-02-23)

* **Bug Fix**: Move all common, SDK-side middleware stack ops into the service client module to prevent cross-module compatibility issues in the future.
* **Dependency Update**: Updated to the latest SDK module versions

# v1.32.0 (2024-02-22)

* **Feature**: Add middleware stack snapshot tests.

# v1.31.2 (2024-02-21)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.31.1 (2024-02-20)

* **Bug Fix**: When sourcing values for a service's `EndpointParameters`, the lack of a configured region (i.e. `options.Region == ""`) will now translate to a `nil` value for `EndpointParameters.Region` instead of a pointer to the empty string `""`. This will result in a much more explicit error when calling an operation instead of an obscure hostname lookup failure.

# v1.31.0 (2024-02-15)

* **Feature**: Adds additional supported instance types.

# v1.30.0 (2024-02-13)

* **Feature**: Bump minimum Go version to 1.20 per our language support policy.
* **Dependency Update**: Updated to the latest SDK module versions

# v1.29.0 (2024-02-06)

* **Feature**: This release adds clear visibility to the customers on the changes that they make on the domain.

# v1.28.0 (2024-01-04)

* **Feature**: This release adds support for new or existing Amazon OpenSearch domains to enable TLS 1.3 or TLS 1.2 with perfect forward secrecy cipher suites for domain endpoints.
* **Dependency Update**: Updated to the latest SDK module versions

# v1.27.1 (2023-12-20)

* No change notes available for this release.

# v1.27.0 (2023-12-14)

* **Feature**: Updating documentation for Amazon OpenSearch Service support for new zero-ETL integration with Amazon S3.

# v1.26.5 (2023-12-08)

* **Bug Fix**: Reinstate presence of default Retryer in functional options, but still respect max attempts set therein.

# v1.26.4 (2023-12-07)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.26.3 (2023-12-06)

* **Bug Fix**: Restore pre-refactor auth behavior where all operations could technically be performed anonymously.

# v1.26.2 (2023-12-01)

* **Bug Fix**: Correct wrapping of errors in authentication workflow.
* **Bug Fix**: Correctly recognize cache-wrapped instances of AnonymousCredentials at client construction.
* **Dependency Update**: Updated to the latest SDK module versions

# v1.26.1 (2023-11-30)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.26.0 (2023-11-29)

* **Feature**: Expose Options() accessor on service clients.
* **Feature**: Launching Amazon OpenSearch Service support for new zero-ETL integration with Amazon S3. Customers can now manage their direct query data sources to Amazon S3 programatically
* **Dependency Update**: Updated to the latest SDK module versions

# v1.25.5 (2023-11-28.2)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.25.4 (2023-11-28)

* **Bug Fix**: Respect setting RetryMaxAttempts in functional options at client construction.

# v1.25.3 (2023-11-20)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.25.2 (2023-11-15)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.25.1 (2023-11-09)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.25.0 (2023-11-01)

* **Feature**: Adds support for configured endpoints via environment variables and the AWS shared configuration file.
* **Dependency Update**: Updated to the latest SDK module versions

# v1.24.0 (2023-10-31)

* **Feature**: **BREAKING CHANGE**: Bump minimum go version to 1.19 per the revised [go version support policy](https://aws.amazon.com/blogs/developer/aws-sdk-for-go-aligns-with-go-release-policy-on-supported-runtimes/).
* **Dependency Update**: Updated to the latest SDK module versions

# v1.23.0 (2023-10-26)

* **Feature**: You can specify ipv4 or dualstack IPAddressType for cluster endpoints. If you specify IPAddressType as dualstack, the new endpoint will be visible under the 'EndpointV2' parameter and will support IPv4 and IPv6 requests. Whereas, the 'Endpoint' will continue to serve IPv4 requests.

# v1.22.0 (2023-10-24)

* **Feature**: **BREAKFIX**: Correct nullability and default value representation of various input fields across a large number of services. Calling code that references one or more of the affected fields will need to update usage accordingly. See [2162](https://github.com/aws/aws-sdk-go-v2/issues/2162).

# v1.21.0 (2023-10-19)

* **Feature**: Added Cluster Administrative options for node restart, opensearch process restart and opensearch dashboard restart for Multi-AZ without standby domains

# v1.20.0 (2023-10-16)

* **Feature**: This release allows customers to list and associate optional plugin packages with compatible Amazon OpenSearch Service clusters for enhanced functionality.

# v1.19.8 (2023-10-12)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.19.7 (2023-10-06)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.19.6 (2023-09-06)

* No change notes available for this release.

# v1.19.5 (2023-08-21)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.19.4 (2023-08-18)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.19.3 (2023-08-17)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.19.2 (2023-08-07)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.19.1 (2023-08-01)

* No change notes available for this release.

# v1.19.0 (2023-07-31)

* **Feature**: Adds support for smithy-modeled endpoint resolution. A new rules-based endpoint resolution will be added to the SDK which will supercede and deprecate existing endpoint resolution. Specifically, EndpointResolver will be deprecated while BaseEndpoint and EndpointResolverV2 will take its place. For more information, please see the Endpoints section in our Developer Guide.
* **Dependency Update**: Updated to the latest SDK module versions

# v1.18.4 (2023-07-28)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.18.3 (2023-07-13)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.18.2 (2023-06-15)

* No change notes available for this release.

# v1.18.1 (2023-06-13)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.18.0 (2023-06-12)

* **Feature**: This release adds support for SkipUnavailable connection property for cross cluster search

# v1.17.0 (2023-05-04)

* **Feature**: DescribeDomainNodes: A new API that provides configuration information for nodes part of the domain

# v1.16.0 (2023-05-03)

* **Feature**: Amazon OpenSearch Service adds the option to deploy a domain across multiple Availability Zones, with each AZ containing a complete copy of data and with nodes in one AZ acting as a standby. This option provides 99.99% availability and consistent performance in the event of infrastructure failure.

# v1.15.5 (2023-04-24)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.15.4 (2023-04-10)

* No change notes available for this release.

# v1.15.3 (2023-04-07)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.15.2 (2023-03-21)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.15.1 (2023-03-10)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.15.0 (2023-02-22)

* **Feature**: This release lets customers configure Off-peak window and software update related properties for a new/existing domain. It enhances the capabilities of StartServiceSoftwareUpdate API; adds 2 new APIs - ListScheduledActions & UpdateScheduledAction; and allows Auto-tune to make use of Off-peak window.
* **Bug Fix**: Prevent nil pointer dereference when retrieving error codes.

# v1.14.3 (2023-02-20)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.14.2 (2023-02-15)

* **Announcement**: When receiving an error response in restJson-based services, an incorrect error type may have been returned based on the content of the response. This has been fixed via PR #2012 tracked in issue #1910.
* **Bug Fix**: Correct error type parsing for restJson services.

# v1.14.1 (2023-02-03)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.14.0 (2023-01-31)

* **Feature**: Amazon OpenSearch Service adds the option for a VPC endpoint connection between two domains when the local domain uses OpenSearch version 1.3 or 2.3. You can now use remote reindex to copy indices from one VPC domain to another without a reverse proxy.

# v1.13.1 (2023-01-23)

* No change notes available for this release.

# v1.13.0 (2023-01-19)

* **Feature**: This release adds the enhanced dry run option, that checks for validation errors that might occur when deploying configuration changes and provides a summary of these errors, if any. The feature will also indicate whether a blue/green deployment will be required to apply a change.

# v1.12.0 (2023-01-05)

* **Feature**: Add `ErrorCodeOverride` field to all error structs (aws/smithy-go#401).

# v1.11.5 (2022-12-15)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.11.4 (2022-12-02)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.11.3 (2022-11-22)

* No change notes available for this release.

# v1.11.2 (2022-11-16)

* No change notes available for this release.

# v1.11.1 (2022-11-10)

* No change notes available for this release.

# v1.11.0 (2022-11-08)

* **Feature**: Amazon OpenSearch Service now offers managed VPC endpoints to connect to your Amazon OpenSearch Service VPC-enabled domain in a Virtual Private Cloud (VPC). This feature allows you to privately access OpenSearch Service domain without using public IPs or requiring traffic to traverse the Internet.

# v1.10.12 (2022-10-24)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.10.11 (2022-10-21)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.10.10 (2022-09-20)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.10.9 (2022-09-14)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.10.8 (2022-09-02)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.10.7 (2022-08-31)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.10.6 (2022-08-30)

* No change notes available for this release.

# v1.10.5 (2022-08-29)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.10.4 (2022-08-11)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.10.3 (2022-08-09)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.10.2 (2022-08-08)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.10.1 (2022-08-01)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.10.0 (2022-07-28)

* **Feature**: This release adds support for gp3 EBS (Elastic Block Store) storage.

# v1.9.9 (2022-07-11)

* No change notes available for this release.

# v1.9.8 (2022-07-05)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.9.7 (2022-06-29)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.9.6 (2022-06-07)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.9.5 (2022-05-17)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.9.4 (2022-04-25)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.9.3 (2022-03-30)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.9.2 (2022-03-24)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.9.1 (2022-03-23)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.9.0 (2022-03-08)

* **Feature**: Updated `github.com/aws/smithy-go` to latest version
* **Dependency Update**: Updated to the latest SDK module versions

# v1.8.0 (2022-02-24)

* **Feature**: API client updated
* **Feature**: Adds RetryMaxAttempts and RetryMod to API client Options. This allows the API clients' default Retryer to be configured from the shared configuration files or environment variables. Adding a new Retry mode of `Adaptive`. `Adaptive` retry mode is an experimental mode, adding client rate limiting when throttles reponses are received from an API. See [retry.AdaptiveMode](https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/aws/retry#AdaptiveMode) for more details, and configuration options.
* **Feature**: Updated `github.com/aws/smithy-go` to latest version
* **Dependency Update**: Updated to the latest SDK module versions

# v1.7.0 (2022-01-14)

* **Feature**: Updated API models
* **Feature**: Updated `github.com/aws/smithy-go` to latest version
* **Dependency Update**: Updated to the latest SDK module versions

# v1.6.0 (2022-01-07)

* **Feature**: Updated `github.com/aws/smithy-go` to latest version
* **Dependency Update**: Updated to the latest SDK module versions

# v1.5.0 (2021-12-21)

* **Feature**: API Paginators now support specifying the initial starting token, and support stopping on empty string tokens.
* **Feature**: Updated to latest service endpoints

# v1.4.1 (2021-12-02)

* **Bug Fix**: Fixes a bug that prevented aws.EndpointResolverWithOptions from being used by the service client. ([#1514](https://github.com/aws/aws-sdk-go-v2/pull/1514))
* **Dependency Update**: Updated to the latest SDK module versions

# v1.4.0 (2021-11-30)

* **Feature**: API client updated

# v1.3.1 (2021-11-19)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.3.0 (2021-11-06)

* **Feature**: The SDK now supports configuration of FIPS and DualStack endpoints using environment variables, shared configuration, or programmatically.
* **Feature**: Updated `github.com/aws/smithy-go` to latest version
* **Feature**: Updated service to latest API model.
* **Dependency Update**: Updated to the latest SDK module versions

# v1.2.0 (2021-10-21)

* **Feature**: Updated  to latest version
* **Dependency Update**: Updated to the latest SDK module versions

# v1.1.1 (2021-10-11)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.1.0 (2021-09-24)

* **Feature**: API client updated

# v1.0.1 (2021-09-17)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.0.0 (2021-09-10)

* **Release**: New AWS service client module
* **Feature**: API client updated


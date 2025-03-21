# v1.15.0 (2024-05-24)

* **Feature**: AWS IoT FleetWise now supports listing vehicles with attributes filter, ListVehicles API is updated to support additional attributes filter.

# v1.14.8 (2024-05-23)

* No change notes available for this release.

# v1.14.7 (2024-05-16)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.14.6 (2024-05-15)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.14.5 (2024-05-08)

* **Bug Fix**: GoDoc improvement

# v1.14.4 (2024-03-29)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.14.3 (2024-03-18)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.14.2 (2024-03-07)

* **Bug Fix**: Remove dependency on go-cmp.
* **Dependency Update**: Updated to the latest SDK module versions

# v1.14.1 (2024-02-23)

* **Bug Fix**: Move all common, SDK-side middleware stack ops into the service client module to prevent cross-module compatibility issues in the future.
* **Dependency Update**: Updated to the latest SDK module versions

# v1.14.0 (2024-02-22)

* **Feature**: Add middleware stack snapshot tests.

# v1.13.3 (2024-02-21)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.13.2 (2024-02-20)

* **Bug Fix**: When sourcing values for a service's `EndpointParameters`, the lack of a configured region (i.e. `options.Region == ""`) will now translate to a `nil` value for `EndpointParameters.Region` instead of a pointer to the empty string `""`. This will result in a much more explicit error when calling an operation instead of an obscure hostname lookup failure.

# v1.13.1 (2024-02-15)

* **Bug Fix**: Correct failure to determine the error type in awsJson services that could occur when errors were modeled with a non-string `code` field.

# v1.13.0 (2024-02-13)

* **Feature**: Bump minimum Go version to 1.20 per our language support policy.
* **Dependency Update**: Updated to the latest SDK module versions

# v1.12.0 (2024-01-16)

* **Feature**: Updated APIs: SignalNodeType query parameter has been added to ListSignalCatalogNodesRequest and ListVehiclesResponse has been extended with attributes field.

# v1.11.0 (2024-01-11)

* **Feature**: The following dataTypes have been removed: CUSTOMER_DECODED_INTERFACE in NetworkInterfaceType; CUSTOMER_DECODED_SIGNAL_INFO_IS_NULL in SignalDecoderFailureReason; CUSTOMER_DECODED_SIGNAL_NETWORK_INTERFACE_INFO_IS_NULL in NetworkInterfaceFailureReason; CUSTOMER_DECODED_SIGNAL in SignalDecoderType

# v1.10.6 (2024-01-04)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.10.5 (2023-12-08)

* **Bug Fix**: Reinstate presence of default Retryer in functional options, but still respect max attempts set therein.

# v1.10.4 (2023-12-07)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.10.3 (2023-12-06)

* **Bug Fix**: Restore pre-refactor auth behavior where all operations could technically be performed anonymously.

# v1.10.2 (2023-12-01)

* **Bug Fix**: Correct wrapping of errors in authentication workflow.
* **Bug Fix**: Correctly recognize cache-wrapped instances of AnonymousCredentials at client construction.
* **Dependency Update**: Updated to the latest SDK module versions

# v1.10.1 (2023-11-30)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.10.0 (2023-11-29)

* **Feature**: Expose Options() accessor on service clients.
* **Dependency Update**: Updated to the latest SDK module versions

# v1.9.2 (2023-11-28.2)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.9.1 (2023-11-28)

* **Bug Fix**: Respect setting RetryMaxAttempts in functional options at client construction.

# v1.9.0 (2023-11-27)

* **Feature**: AWS IoT FleetWise introduces new APIs for vision system data, such as data collected from cameras, radars, and lidars. You can now model and decode complex data types.

# v1.8.3 (2023-11-20)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.8.2 (2023-11-15)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.8.1 (2023-11-09)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.8.0 (2023-11-01)

* **Feature**: Adds support for configured endpoints via environment variables and the AWS shared configuration file.
* **Dependency Update**: Updated to the latest SDK module versions

# v1.7.0 (2023-10-31)

* **Feature**: **BREAKING CHANGE**: Bump minimum go version to 1.19 per the revised [go version support policy](https://aws.amazon.com/blogs/developer/aws-sdk-for-go-aligns-with-go-release-policy-on-supported-runtimes/).
* **Dependency Update**: Updated to the latest SDK module versions

# v1.6.2 (2023-10-12)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.6.1 (2023-10-06)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.6.0 (2023-09-28)

* **Feature**: AWS IoT FleetWise now supports encryption through a customer managed AWS KMS key. The PutEncryptionConfiguration and GetEncryptionConfiguration APIs were added.

# v1.5.5 (2023-08-21)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.5.4 (2023-08-18)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.5.3 (2023-08-17)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.5.2 (2023-08-07)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.5.1 (2023-08-01)

* No change notes available for this release.

# v1.5.0 (2023-07-31)

* **Feature**: Adds support for smithy-modeled endpoint resolution. A new rules-based endpoint resolution will be added to the SDK which will supercede and deprecate existing endpoint resolution. Specifically, EndpointResolver will be deprecated while BaseEndpoint and EndpointResolverV2 will take its place. For more information, please see the Endpoints section in our Developer Guide.
* **Dependency Update**: Updated to the latest SDK module versions

# v1.4.4 (2023-07-28)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.4.3 (2023-07-13)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.4.2 (2023-06-15)

* No change notes available for this release.

# v1.4.1 (2023-06-13)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.4.0 (2023-05-30)

* **Feature**: Campaigns now support selecting Timestream or S3 as the data destination, Signal catalogs now support "Deprecation" keyword released in VSS v2.1 and "Comment" keyword released in VSS v3.0

# v1.3.10 (2023-05-04)

* No change notes available for this release.

# v1.3.9 (2023-04-24)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.3.8 (2023-04-10)

* No change notes available for this release.

# v1.3.7 (2023-04-07)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.3.6 (2023-03-21)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.3.5 (2023-03-10)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.3.4 (2023-02-22)

* **Bug Fix**: Prevent nil pointer dereference when retrieving error codes.

# v1.3.3 (2023-02-20)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.3.2 (2023-02-15)

* **Announcement**: When receiving an error response in restJson-based services, an incorrect error type may have been returned based on the content of the response. This has been fixed via PR #2012 tracked in issue #1910.
* **Bug Fix**: Correct error type parsing for restJson services.

# v1.3.1 (2023-02-03)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.3.0 (2023-01-05)

* **Feature**: Add `ErrorCodeOverride` field to all error structs (aws/smithy-go#401).

# v1.2.1 (2022-12-30)

* **Documentation**: Update documentation - correct the epoch constant value of default value for expiryTime field in CreateCampaign request.

# v1.2.0 (2022-12-16)

* **Feature**: Updated error handling for empty resource names in "UpdateSignalCatalog" and "GetModelManifest" operations.

# v1.1.1 (2022-12-15)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.1.0 (2022-12-09)

* **Feature**: Deprecated assignedValue property for actuators and attributes.  Added a message to invalid nodes and invalid decoder manifest exceptions.

# v1.0.4 (2022-12-02)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.0.3 (2022-10-24)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.0.2 (2022-10-21)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.0.1 (2022-10-13)

* **Documentation**: Documentation update for AWS IoT FleetWise

# v1.0.0 (2022-09-26)

* **Release**: New AWS service client module
* **Feature**: General availability (GA) for AWS IoT Fleetwise. It adds AWS IoT Fleetwise to AWS SDK. For more information, see https://docs.aws.amazon.com/iot-fleetwise/latest/APIReference/Welcome.html.


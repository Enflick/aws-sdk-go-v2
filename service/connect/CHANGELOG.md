# v1.100.0 (2024-05-30)

* **Feature**: Adding associatedQueueIds as a SearchCriteria and response field to the SearchRoutingProfiles API

# v1.99.0 (2024-05-29)

* **Feature**: This release includes changes to DescribeContact API's response by including ConnectedToSystemTimestamp, RoutingCriteria, Customer, Campaign, AnsweringMachineDetectionStatus, CustomerVoiceActivity, QualityMetrics, DisconnectDetails, and SegmentAttributes information from a contact in Amazon Connect.

# v1.98.3 (2024-05-23)

* No change notes available for this release.

# v1.98.2 (2024-05-16)

* **Documentation**: Adding Contact Flow metrics to the GetMetricDataV2 API
* **Dependency Update**: Updated to the latest SDK module versions

# v1.98.1 (2024-05-15)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.98.0 (2024-05-14)

* **Feature**: Amazon Connect provides enhanced search capabilities for flows & flow modules on the Connect admin website and programmatically using APIs. You can search for flows and flow modules by name, description, type, status, and tags, to filter and identify a specific flow in your Connect instances.

# v1.97.1 (2024-05-08)

* **Bug Fix**: GoDoc improvement

# v1.97.0 (2024-05-03)

* **Feature**: This release adds 5 new APIs for managing attachments: StartAttachedFileUpload, CompleteAttachedFileUpload, GetAttachedFile, BatchGetAttachedFileMetadata, DeleteAttachedFile. These APIs can be used to programmatically upload and download attachments to Connect resources, like cases.

# v1.96.0 (2024-04-10)

* **Feature**: This release adds new Submit Auto Evaluation Action for Amazon Connect Rules.

# v1.95.1 (2024-03-29)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.95.0 (2024-03-20)

* **Feature**: This release updates the *InstanceStorageConfig APIs to support a new ResourceType: REAL_TIME_CONTACT_ANALYSIS_CHAT_SEGMENTS. Use this resource type to enable streaming for real-time analysis of chat contacts and to associate a Kinesis stream where real-time analysis chat segments will be published.

# v1.94.1 (2024-03-18)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.94.0 (2024-03-15)

* **Feature**: This release adds Hierarchy based Access Control fields to Security Profile public APIs and adds support for UserAttributeFilter to SearchUsers API.

# v1.93.0 (2024-03-12)

* **Feature**: This release increases MaxResults limit to 500 in request for SearchUsers, SearchQueues and SearchRoutingProfiles APIs of Amazon Connect.

# v1.92.2 (2024-03-07)

* **Bug Fix**: Remove dependency on go-cmp.
* **Dependency Update**: Updated to the latest SDK module versions

# v1.92.1 (2024-02-23)

* **Bug Fix**: Move all common, SDK-side middleware stack ops into the service client module to prevent cross-module compatibility issues in the future.
* **Dependency Update**: Updated to the latest SDK module versions

# v1.92.0 (2024-02-22)

* **Feature**: Add middleware stack snapshot tests.

# v1.91.2 (2024-02-21)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.91.1 (2024-02-20)

* **Bug Fix**: When sourcing values for a service's `EndpointParameters`, the lack of a configured region (i.e. `options.Region == ""`) will now translate to a `nil` value for `EndpointParameters.Region` instead of a pointer to the empty string `""`. This will result in a much more explicit error when calling an operation instead of an obscure hostname lookup failure.

# v1.91.0 (2024-02-13)

* **Feature**: Bump minimum Go version to 1.20 per our language support policy.
* **Dependency Update**: Updated to the latest SDK module versions

# v1.90.0 (2024-01-18)

* **Feature**: GetMetricDataV2 now supports 3 groupings

# v1.89.0 (2024-01-12)

* **Feature**: Supervisor Barge for Chat is now supported through the MonitorContact API.

# v1.88.1 (2024-01-04)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.88.0 (2024-01-03)

* **Feature**: Amazon Connect, Contact Lens Evaluation API increase evaluation notes max length to 3072.

# v1.87.0 (2023-12-21)

* **Feature**: Adds APIs to manage User Proficiencies and Predefined Attributes. Enhances StartOutboundVoiceContact API input. Introduces SearchContacts API. Enhances DescribeContact API. Adds an API to update Routing Attributes in QueuePriority and QueueTimeAdjustmentSeconds.

# v1.86.0 (2023-12-15)

* **Feature**: Adds relatedContactId field to StartOutboundVoiceContact API input. Introduces PauseContact API and ResumeContact API for Task contacts. Adds pause duration, number of pauses, timestamps for last paused and resumed events to DescribeContact API response. Adds new Rule type and new Rule action.

# v1.85.0 (2023-12-14)

* **Feature**: This release adds support for more granular billing using tags (key:value pairs)

# v1.84.2 (2023-12-08)

* **Bug Fix**: Reinstate presence of default Retryer in functional options, but still respect max attempts set therein.

# v1.84.1 (2023-12-07)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.84.0 (2023-12-06)

* **Feature**: Releasing Tagging Support for Instance Management APIS
* **Bug Fix**: Restore pre-refactor auth behavior where all operations could technically be performed anonymously.

# v1.83.2 (2023-12-01)

* **Bug Fix**: Correct wrapping of errors in authentication workflow.
* **Bug Fix**: Correctly recognize cache-wrapped instances of AnonymousCredentials at client construction.
* **Dependency Update**: Updated to the latest SDK module versions

# v1.83.1 (2023-11-30)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.83.0 (2023-11-29)

* **Feature**: Expose Options() accessor on service clients.
* **Dependency Update**: Updated to the latest SDK module versions

# v1.82.0 (2023-11-28.2)

* **Feature**: Added support for following capabilities: Amazon Connect's in-app, web, and video calling. Two-way SMS integrations. Contact Lens real-time chat analytics feature. Amazon Connect Analytics Datalake capability. Capability to configure real time chat rules.
* **Dependency Update**: Updated to the latest SDK module versions

# v1.81.2 (2023-11-28)

* **Bug Fix**: Respect setting RetryMaxAttempts in functional options at client construction.

# v1.81.1 (2023-11-20)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.81.0 (2023-11-17)

* **Feature**: This release adds WISDOM_QUICK_RESPONSES as new IntegrationType of Connect IntegrationAssociation resource and bug fixes.

# v1.80.1 (2023-11-15)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.80.0 (2023-11-14)

* **Feature**: Introducing SegmentAttributes parameter for StartChatContact API

# v1.79.0 (2023-11-09.2)

* **Feature**: This release adds the ability to integrate customer lambda functions with Connect attachments for scanning and updates the ListIntegrationAssociations API to support filtering on IntegrationArn.

# v1.78.1 (2023-11-09)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.78.0 (2023-11-08)

* **Feature**: This release clarifies in our public documentation that InstanceId is a requirement for SearchUsers API requests.

# v1.77.0 (2023-11-06)

* **Feature**: Added new API that allows Amazon Connect Outbound Campaigns to create contacts in Amazon Connect when ingesting your dial requests.

# v1.76.0 (2023-11-03)

* **Feature**: Amazon Connect Chat introduces Create Persistent Contact Association API, allowing customers to choose when to resume previous conversations from previous chats, eliminating the need to repeat themselves and allowing agents to provide personalized service with access to entire conversation history.

# v1.75.0 (2023-11-02)

* **Feature**: GetMetricDataV2 API: Update to include new metrics PERCENT_NON_TALK_TIME, PERCENT_TALK_TIME, PERCENT_TALK_TIME_AGENT, PERCENT_TALK_TIME_CUSTOMER

# v1.74.0 (2023-11-01)

* **Feature**: Adds support for configured endpoints via environment variables and the AWS shared configuration file.
* **Feature**: Adds the BatchGetFlowAssociation API which returns flow associations (flow-resource) corresponding to the list of resourceArns supplied in the request. This release also adds IsDefault, LastModifiedRegion and LastModifiedTime fields to the responses of several Describe and List APIs.
* **Dependency Update**: Updated to the latest SDK module versions

# v1.73.0 (2023-10-31)

* **Feature**: **BREAKING CHANGE**: Bump minimum go version to 1.19 per the revised [go version support policy](https://aws.amazon.com/blogs/developer/aws-sdk-for-go-aligns-with-go-release-policy-on-supported-runtimes/).
* **Dependency Update**: Updated to the latest SDK module versions

# v1.72.0 (2023-10-30)

* **Feature**: This release adds InstanceId field for phone number APIs.

# v1.71.0 (2023-10-24)

* **Feature**: **BREAKFIX**: Correct nullability and default value representation of various input fields across a large number of services. Calling code that references one or more of the affected fields will need to update usage accordingly. See [2162](https://github.com/aws/aws-sdk-go-v2/issues/2162).

# v1.70.0 (2023-10-20)

* **Feature**: This release adds support for updating phone number metadata, such as phone number description.

# v1.69.2 (2023-10-12)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.69.1 (2023-10-06)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.69.0 (2023-10-03)

* **Feature**: GetMetricDataV2 API: Update to include new metrics CONTACTS_RESOLVED_IN_X , AVG_HOLD_TIME_ALL_CONTACTS , AVG_RESOLUTION_TIME , ABANDONMENT_RATE , AGENT_NON_RESPONSE_WITHOUT_CUSTOMER_ABANDONS with added features: Interval Period, TimeZone, Negate MetricFilters, Extended date time range.

# v1.68.0 (2023-09-26)

* **Feature**: This release updates a set of Amazon Connect APIs that provides the ability to integrate third party applications in the Amazon Connect agent workspace.

# v1.67.0 (2023-09-15)

* **Feature**: New rule type (OnMetricDataUpdate) has been added

# v1.66.0 (2023-09-01)

* **Feature**: Amazon Connect adds the ability to read, create, update, delete, and list view resources, and adds the ability to read, create, delete, and list view versions.

# v1.65.3 (2023-08-21)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.65.2 (2023-08-18)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.65.1 (2023-08-17)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.65.0 (2023-08-10)

* **Feature**: This release adds APIs to provision agents that are global / available in multiple AWS regions and distribute them across these regions by percentage.

# v1.64.1 (2023-08-07)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.64.0 (2023-08-04)

* **Feature**: Added a new API UpdateRoutingProfileAgentAvailabilityTimer to update agent availability timer of a routing profile.

# v1.63.1 (2023-08-01)

* No change notes available for this release.

# v1.63.0 (2023-07-31)

* **Feature**: Adds support for smithy-modeled endpoint resolution. A new rules-based endpoint resolution will be added to the SDK which will supercede and deprecate existing endpoint resolution. Specifically, EndpointResolver will be deprecated while BaseEndpoint and EndpointResolverV2 will take its place. For more information, please see the Endpoints section in our Developer Guide.
* **Dependency Update**: Updated to the latest SDK module versions

# v1.62.0 (2023-07-28.2)

* **Feature**: This release adds support for new number types.

# v1.61.2 (2023-07-28)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.61.1 (2023-07-18)

* **Documentation**: GetMetricDataV2 API: Update to include Contact Lens Conversational Analytics Metrics

# v1.61.0 (2023-07-13)

* **Feature**: Add support for deleting Queues and Routing Profiles.
* **Dependency Update**: Updated to the latest SDK module versions

# v1.60.1 (2023-07-05)

* **Documentation**: GetMetricDataV2 API: Channels filters do not count towards overall limitation of 100 filter values.

# v1.60.0 (2023-06-26)

* **Feature**: This release provides a way to search for existing tags within an instance. Before tagging a resource, ensure consistency by searching for pre-existing key:value pairs.

# v1.59.0 (2023-06-16)

* **Feature**: Updates the *InstanceStorageConfig APIs to support a new ResourceType: SCREEN_RECORDINGS to enable screen recording and specify the storage configurations for publishing the recordings. Also updates DescribeInstance and ListInstances APIs to include InstanceAccessUrl attribute in the API response.

# v1.58.2 (2023-06-15)

* No change notes available for this release.

# v1.58.1 (2023-06-13)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.58.0 (2023-06-09)

* **Feature**: This release adds search APIs for Prompts, Quick Connects and Hours of Operations, which can be used to search for those resources within a Connect Instance.

# v1.57.1 (2023-06-06)

* **Documentation**: GetMetricDataV2 API is now available in AWS GovCloud(US) region.

# v1.57.0 (2023-05-26)

* **Feature**: Documentation update for a new Initiation Method value in DescribeContact API

# v1.56.0 (2023-05-24)

* **Feature**: Amazon Connect Evaluation Capabilities: validation improvements

# v1.55.0 (2023-05-18)

* **Feature**: You can programmatically create and manage prompts using APIs, for example, to extract prompts stored within Amazon Connect and add them to your Amazon S3 bucket. AWS CloudTrail, AWS CloudFormation and tagging are supported.

# v1.54.2 (2023-05-11)

* **Documentation**: This release updates GetMetricDataV2 API, to support metric data up-to last 35 days

# v1.54.1 (2023-05-04)

* **Documentation**: Remove unused InvalidParameterException from CreateParticipant API

# v1.54.0 (2023-05-02)

* **Feature**: Amazon Connect Service Rules API update: Added OnContactEvaluationSubmit event source to support user configuring evaluation form rules.

# v1.53.0 (2023-04-25)

* **Feature**: Amazon Connect, Contact Lens Evaluation API release including ability to manage forms and to submit contact evaluations.

# v1.52.1 (2023-04-24)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.52.0 (2023-04-21)

* **Feature**: This release adds a new API CreateParticipant. For Amazon Connect Chat, you can use this new API to customize chat flow experiences.

# v1.51.0 (2023-04-10)

* **Feature**: This release adds the ability to configure an agent's routing profile to receive contacts from multiple channels at the same time via extending the UpdateRoutingProfileConcurrency, CreateRoutingProfile and DescribeRoutingProfile APIs.

# v1.50.1 (2023-04-07)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.50.0 (2023-03-27)

* **Feature**: This release introduces support for RelatedContactId in the StartChatContact API. Interactive message and interactive message response have been added to the list of supported message content types for this API as well.

# v1.49.2 (2023-03-21)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.49.1 (2023-03-10)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.49.0 (2023-03-09)

* **Feature**: This release adds a new API, GetMetricDataV2, which returns metric data for Amazon Connect.

# v1.48.0 (2023-02-24)

* **Feature**: StartTaskContact API now supports linked task creation with a new optional RelatedContactId parameter

# v1.47.1 (2023-02-22)

* **Bug Fix**: Prevent nil pointer dereference when retrieving error codes.

# v1.47.0 (2023-02-20)

* **Feature**: Reasons for failed diff has been approved by SDK Reviewer
* **Dependency Update**: Updated to the latest SDK module versions

# v1.46.1 (2023-02-15)

* **Announcement**: When receiving an error response in restJson-based services, an incorrect error type may have been returned based on the content of the response. This has been fixed via PR #2012 tracked in issue #1910.
* **Bug Fix**: Correct error type parsing for restJson services.

# v1.46.0 (2023-02-10)

* **Feature**: This update provides the Wisdom session ARN for contacts enabled for Wisdom in the chat channel.

# v1.45.1 (2023-02-03)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.45.0 (2023-01-19)

* **Feature**: Amazon Connect Chat introduces Persistent Chat, allowing customers to resume previous conversations with context and transcripts carried over from previous chats, eliminating the need to repeat themselves and allowing agents to provide personalized service with access to entire conversation history.

# v1.44.0 (2023-01-13)

* **Feature**: This release updates the responses of UpdateContactFlowContent, UpdateContactFlowMetadata, UpdateContactFlowName and DeleteContactFlow API with empty responses.

# v1.43.0 (2023-01-05)

* **Feature**: Add `ErrorCodeOverride` field to all error structs (aws/smithy-go#401).
* **Feature**: Documentation update for a new Initiation Method value in DescribeContact API

# v1.42.0 (2022-12-23)

* **Feature**: Support for Routing Profile filter, SortCriteria, and grouping by Routing Profiles for GetCurrentMetricData API. Support for RoutingProfiles, UserHierarchyGroups, and Agents as filters, NextStatus and AgentStatusName for GetCurrentUserData. Adds ApproximateTotalCount to both APIs.

# v1.41.0 (2022-12-22)

* **Feature**: Amazon Connect Chat introduces the Idle Participant/Autodisconnect feature, which allows users to set timeouts relating to the activity of chat participants, using the new UpdateParticipantRoleConfig API.

# v1.40.0 (2022-12-15)

* **Feature**: Added support for "English - New Zealand" and "English - South African" to be used with Amazon Connect Custom Vocabulary APIs.
* **Dependency Update**: Updated to the latest SDK module versions

# v1.39.0 (2022-12-06)

* **Feature**: This release provides APIs that enable you to programmatically manage rules for Contact Lens conversational analytics and third party applications. For more information, see   https://docs.aws.amazon.com/connect/latest/APIReference/rules-api.html

# v1.38.1 (2022-12-02)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.38.0 (2022-11-18)

* **Feature**: Added AllowedAccessControlTags and TagRestrictedResource for Tag Based Access Control on Amazon Connect Webpage

# v1.37.0 (2022-11-16)

* **Feature**: This release adds a new MonitorContact API for initiating monitoring of ongoing Voice and Chat contacts.

# v1.36.0 (2022-11-15)

* **Feature**: This release updates the APIs: UpdateInstanceAttribute, DescribeInstanceAttribute, and ListInstanceAttributes. You can use it to programmatically enable/disable enhanced contact monitoring using attribute type ENHANCED_CONTACT_MONITORING on the specified Amazon Connect instance.

# v1.35.0 (2022-11-09)

* **Feature**: This release adds new fields SignInUrl, UserArn, and UserId to GetFederationToken response payload.

# v1.34.0 (2022-10-31)

* **Feature**: Amazon connect now support a new API DismissUserContact to dismiss or remove terminated contacts in Agent CCP

# v1.33.2 (2022-10-24)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.33.1 (2022-10-21)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.33.0 (2022-10-19)

* **Feature**: This release adds API support for managing phone numbers that can be used across multiple AWS regions through telephony traffic distribution.

# v1.32.0 (2022-10-13)

* **Feature**: This release adds support for a secondary email and a mobile number for Amazon Connect instance users.

# v1.31.0 (2022-10-04)

* **Feature**: Updated the CreateIntegrationAssociation API to support the CASES_DOMAIN IntegrationType.

# v1.30.1 (2022-09-20)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.30.0 (2022-09-14)

* **Feature**: Fixed a bug in the API client generation which caused some operation parameters to be incorrectly generated as value types instead of pointer types. The service API always required these affected parameters to be nilable. This fixes the SDK client to match the expectations of the the service API.
* **Dependency Update**: Updated to the latest SDK module versions

# v1.29.0 (2022-09-02)

* **Feature**: This release adds search APIs for Routing Profiles and Queues, which can be used to search for those resources within a Connect Instance.
* **Dependency Update**: Updated to the latest SDK module versions

# v1.28.2 (2022-08-31)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.28.1 (2022-08-29)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.28.0 (2022-08-19)

* **Feature**: This release adds SearchSecurityProfiles API which can be used to search for Security Profile resources within a Connect Instance.

# v1.27.6 (2022-08-11)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.27.5 (2022-08-09)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.27.4 (2022-08-08)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.27.3 (2022-08-01)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.27.2 (2022-07-05)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.27.1 (2022-06-29)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.27.0 (2022-06-17)

* **Feature**: This release updates these APIs: UpdateInstanceAttribute, DescribeInstanceAttribute and ListInstanceAttributes. You can use it to programmatically enable/disable High volume outbound communications using attribute type HIGH_VOLUME_OUTBOUND on the specified Amazon Connect instance.

# v1.26.1 (2022-06-07)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.26.0 (2022-06-06)

* **Feature**: This release adds a new API, GetCurrentUserData, which returns real-time details about users' current activity.

# v1.25.0 (2022-06-02)

* **Feature**: This release adds the following features: 1) New APIs to manage (create, list, update) task template resources, 2) Updates to startTaskContact API to support task templates, and 3) new TransferContact API to programmatically transfer in-progress tasks via a contact flow.

# v1.24.1 (2022-05-17)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.24.0 (2022-04-28)

* **Feature**: This release introduces an API for changing the current agent status of a user in Connect.

# v1.23.0 (2022-04-25)

* **Feature**: This release adds SearchUsers API which can be used to search for users with a Connect Instance
* **Dependency Update**: Updated to the latest SDK module versions

# v1.22.0 (2022-04-20)

* **Feature**: This release adds APIs to search, claim, release, list, update, and describe phone numbers. You can also use them to associate and disassociate contact flows to phone numbers.

# v1.21.0 (2022-04-01)

* **Feature**: This release updates these APIs: UpdateInstanceAttribute, DescribeInstanceAttribute and ListInstanceAttributes. You can use it to programmatically enable/disable multi-party conferencing using attribute type MULTI_PARTY_CONFERENCING on the specified Amazon Connect instance.

# v1.20.3 (2022-03-30)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.20.2 (2022-03-24)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.20.1 (2022-03-23)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.20.0 (2022-03-11)

* **Feature**: This release adds support for enabling Rich Messaging when starting a new chat session via the StartChatContact API. Rich Messaging enables the following formatting options: bold, italics, hyperlinks, bulleted lists, and numbered lists.

# v1.19.0 (2022-03-08)

* **Feature**: Updated `github.com/aws/smithy-go` to latest version
* **Feature**: Updated service client model to latest release.
* **Dependency Update**: Updated to the latest SDK module versions

# v1.18.0 (2022-02-24)

* **Feature**: API client updated
* **Feature**: Adds RetryMaxAttempts and RetryMod to API client Options. This allows the API clients' default Retryer to be configured from the shared configuration files or environment variables. Adding a new Retry mode of `Adaptive`. `Adaptive` retry mode is an experimental mode, adding client rate limiting when throttles reponses are received from an API. See [retry.AdaptiveMode](https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/aws/retry#AdaptiveMode) for more details, and configuration options.
* **Feature**: Updated `github.com/aws/smithy-go` to latest version
* **Dependency Update**: Updated to the latest SDK module versions

# v1.17.0 (2022-01-28)

* **Feature**: Updated to latest API model.

# v1.16.0 (2022-01-14)

* **Feature**: Updated `github.com/aws/smithy-go` to latest version
* **Dependency Update**: Updated to the latest SDK module versions

# v1.15.0 (2022-01-07)

* **Feature**: Updated `github.com/aws/smithy-go` to latest version
* **Dependency Update**: Updated to the latest SDK module versions

# v1.14.0 (2021-12-21)

* **Feature**: API Paginators now support specifying the initial starting token, and support stopping on empty string tokens.

# v1.13.1 (2021-12-02)

* **Bug Fix**: Fixes a bug that prevented aws.EndpointResolverWithOptions from being used by the service client. ([#1514](https://github.com/aws/aws-sdk-go-v2/pull/1514))
* **Dependency Update**: Updated to the latest SDK module versions

# v1.13.0 (2021-11-30)

* **Feature**: API client updated

# v1.12.0 (2021-11-19)

* **Feature**: API client updated
* **Dependency Update**: Updated to the latest SDK module versions

# v1.11.0 (2021-11-12)

* **Feature**: Updated service to latest API model.

# v1.10.0 (2021-11-06)

* **Feature**: The SDK now supports configuration of FIPS and DualStack endpoints using environment variables, shared configuration, or programmatically.
* **Feature**: Updated `github.com/aws/smithy-go` to latest version
* **Feature**: Updated service to latest API model.
* **Dependency Update**: Updated to the latest SDK module versions

# v1.9.0 (2021-10-21)

* **Feature**: Updated  to latest version
* **Dependency Update**: Updated to the latest SDK module versions

# v1.8.1 (2021-10-11)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.8.0 (2021-09-30)

* **Feature**: API client updated

# v1.7.1 (2021-09-17)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.7.0 (2021-08-27)

* **Feature**: Updated `github.com/aws/smithy-go` to latest version
* **Dependency Update**: Updated to the latest SDK module versions

# v1.6.1 (2021-08-19)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.6.0 (2021-08-12)

* **Feature**: API client updated

# v1.5.2 (2021-08-04)

* **Dependency Update**: Updated `github.com/aws/smithy-go` to latest version.
* **Dependency Update**: Updated to the latest SDK module versions

# v1.5.1 (2021-07-15)

* **Dependency Update**: Updated `github.com/aws/smithy-go` to latest version
* **Dependency Update**: Updated to the latest SDK module versions

# v1.5.0 (2021-06-25)

* **Feature**: API client updated
* **Feature**: Updated `github.com/aws/smithy-go` to latest version
* **Dependency Update**: Updated to the latest SDK module versions

# v1.4.1 (2021-05-20)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.4.0 (2021-05-14)

* **Feature**: Constant has been added to modules to enable runtime version inspection for reporting.
* **Feature**: Updated to latest service API model.
* **Dependency Update**: Updated to the latest SDK module versions


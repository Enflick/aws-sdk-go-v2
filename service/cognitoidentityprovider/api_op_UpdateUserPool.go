// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This action might generate an SMS text message. Starting June 1, 2021, US
// telecom carriers require you to register an origination phone number before you
// can send SMS messages to US phone numbers. If you use SMS text messages in
// Amazon Cognito, you must register a phone number with [Amazon Pinpoint]. Amazon Cognito uses the
// registered number automatically. Otherwise, Amazon Cognito users who must
// receive SMS messages might not be able to sign up, activate their accounts, or
// sign in.
//
// If you have never used SMS text messages with Amazon Cognito or any other
// Amazon Web Service, Amazon Simple Notification Service might place your account
// in the SMS sandbox. In [sandbox mode], you can send messages only to verified phone numbers.
// After you test your app while in the sandbox environment, you can move out of
// the sandbox and into production. For more information, see [SMS message settings for Amazon Cognito user pools]in the Amazon
// Cognito Developer Guide.
//
// Updates the specified user pool with the specified attributes. You can get a
// list of the current user pool settings using [DescribeUserPool].
//
// If you don't provide a value for an attribute, Amazon Cognito sets it to its
// default value.
//
// Amazon Cognito evaluates Identity and Access Management (IAM) policies in
// requests for this API operation. For this operation, you must use IAM
// credentials to authorize requests, and you must grant yourself the corresponding
// IAM permission in a policy.
//
// # Learn more
//
// [Signing Amazon Web Services API Requests]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]
//
// [SMS message settings for Amazon Cognito user pools]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-sms-settings.html
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [DescribeUserPool]: https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_DescribeUserPool.html
// [sandbox mode]: https://docs.aws.amazon.com/sns/latest/dg/sns-sms-sandbox.html
// [Signing Amazon Web Services API Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html
// [Amazon Pinpoint]: https://console.aws.amazon.com/pinpoint/home/
func (c *Client) UpdateUserPool(ctx context.Context, params *UpdateUserPoolInput, optFns ...func(*Options)) (*UpdateUserPoolOutput, error) {
	if params == nil {
		params = &UpdateUserPoolInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateUserPool", params, optFns, c.addOperationUpdateUserPoolMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateUserPoolOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the request to update the user pool.
type UpdateUserPoolInput struct {

	// The user pool ID for the user pool you want to update.
	//
	// This member is required.
	UserPoolId *string

	// The available verified method a user can use to recover their password when
	// they call ForgotPassword . You can use this setting to define a preferred method
	// when a user has more than one method available. With this setting, SMS doesn't
	// qualify for a valid password recovery mechanism if the user also has SMS
	// multi-factor authentication (MFA) activated. In the absence of this setting,
	// Amazon Cognito uses the legacy behavior to determine the recovery method where
	// SMS is preferred through email.
	AccountRecoverySetting *types.AccountRecoverySettingType

	// The configuration for AdminCreateUser requests.
	AdminCreateUserConfig *types.AdminCreateUserConfigType

	// The attributes that are automatically verified when Amazon Cognito requests to
	// update user pools.
	AutoVerifiedAttributes []types.VerifiedAttributeType

	// When active, DeletionProtection prevents accidental deletion of your user pool.
	// Before you can delete a user pool that you have protected against deletion, you
	// must deactivate this feature.
	//
	// When you try to delete a protected user pool in a DeleteUserPool API request,
	// Amazon Cognito returns an InvalidParameterException error. To delete a
	// protected user pool, send a new DeleteUserPool request after you deactivate
	// deletion protection in an UpdateUserPool API request.
	DeletionProtection types.DeletionProtectionType

	// The device-remembering configuration for a user pool. A null value indicates
	// that you have deactivated device remembering in your user pool.
	//
	// When you provide a value for any DeviceConfiguration field, you activate the
	// Amazon Cognito device-remembering feature.
	DeviceConfiguration *types.DeviceConfigurationType

	// The email configuration of your user pool. The email configuration type sets
	// your preferred sending method, Amazon Web Services Region, and sender for email
	// invitation and verification messages from your user pool.
	EmailConfiguration *types.EmailConfigurationType

	// This parameter is no longer used. See [VerificationMessageTemplateType].
	//
	// [VerificationMessageTemplateType]: https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_VerificationMessageTemplateType.html
	EmailVerificationMessage *string

	// This parameter is no longer used. See [VerificationMessageTemplateType].
	//
	// [VerificationMessageTemplateType]: https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_VerificationMessageTemplateType.html
	EmailVerificationSubject *string

	// The Lambda configuration information from the request to update the user pool.
	LambdaConfig *types.LambdaConfigType

	// Possible values include:
	//
	//   - OFF - MFA tokens aren't required and can't be specified during user
	//   registration.
	//
	//   - ON - MFA tokens are required for all user registrations. You can only
	//   specify ON when you're initially creating a user pool. You can use the [SetUserPoolMfaConfig]API
	//   operation to turn MFA "ON" for existing user pools.
	//
	//   - OPTIONAL - Users have the option when registering to create an MFA token.
	//
	// [SetUserPoolMfaConfig]: https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_SetUserPoolMfaConfig.html
	MfaConfiguration types.UserPoolMfaType

	// A container with the policies you want to update in a user pool.
	Policies *types.UserPoolPolicyType

	// The contents of the SMS authentication message.
	SmsAuthenticationMessage *string

	// The SMS configuration with the settings that your Amazon Cognito user pool must
	// use to send an SMS message from your Amazon Web Services account through Amazon
	// Simple Notification Service. To send SMS messages with Amazon SNS in the Amazon
	// Web Services Region that you want, the Amazon Cognito user pool uses an Identity
	// and Access Management (IAM) role in your Amazon Web Services account.
	SmsConfiguration *types.SmsConfigurationType

	// This parameter is no longer used. See [VerificationMessageTemplateType].
	//
	// [VerificationMessageTemplateType]: https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_VerificationMessageTemplateType.html
	SmsVerificationMessage *string

	// The settings for updates to user attributes. These settings include the
	// property AttributesRequireVerificationBeforeUpdate , a user-pool setting that
	// tells Amazon Cognito how to handle changes to the value of your users' email
	// address and phone number attributes. For more information, see [Verifying updates to email addresses and phone numbers].
	//
	// [Verifying updates to email addresses and phone numbers]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-settings-email-phone-verification.html#user-pool-settings-verifications-verify-attribute-updates
	UserAttributeUpdateSettings *types.UserAttributeUpdateSettingsType

	// User pool add-ons. Contains settings for activation of advanced security
	// features. To log user security information but take no action, set to AUDIT . To
	// configure automatic security responses to risky traffic to your user pool, set
	// to ENFORCED .
	//
	// For more information, see [Adding advanced security to a user pool].
	//
	// [Adding advanced security to a user pool]: https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pool-settings-advanced-security.html
	UserPoolAddOns *types.UserPoolAddOnsType

	// The tag keys and values to assign to the user pool. A tag is a label that you
	// can use to categorize and manage user pools in different ways, such as by
	// purpose, owner, environment, or other criteria.
	UserPoolTags map[string]string

	// The template for verification messages.
	VerificationMessageTemplate *types.VerificationMessageTemplateType

	noSmithyDocumentSerde
}

// Represents the response from the server when you make a request to update the
// user pool.
type UpdateUserPoolOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateUserPoolMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateUserPool{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateUserPool{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateUserPool"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpUpdateUserPoolValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateUserPool(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opUpdateUserPool(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateUserPool",
	}
}

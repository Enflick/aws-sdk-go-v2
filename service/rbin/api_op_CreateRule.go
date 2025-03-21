// Code generated by smithy-go-codegen DO NOT EDIT.

package rbin

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/rbin/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a Recycle Bin retention rule. For more information, see [Create Recycle Bin retention rules] in the Amazon
// Elastic Compute Cloud User Guide.
//
// [Create Recycle Bin retention rules]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/recycle-bin-working-with-rules.html#recycle-bin-create-rule
func (c *Client) CreateRule(ctx context.Context, params *CreateRuleInput, optFns ...func(*Options)) (*CreateRuleOutput, error) {
	if params == nil {
		params = &CreateRuleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateRule", params, optFns, c.addOperationCreateRuleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateRuleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateRuleInput struct {

	// The resource type to be retained by the retention rule. Currently, only Amazon
	// EBS snapshots and EBS-backed AMIs are supported. To retain snapshots, specify
	// EBS_SNAPSHOT . To retain EBS-backed AMIs, specify EC2_IMAGE .
	//
	// This member is required.
	ResourceType types.ResourceType

	// Information about the retention period for which the retention rule is to
	// retain resources.
	//
	// This member is required.
	RetentionPeriod *types.RetentionPeriod

	// The retention rule description.
	Description *string

	// Information about the retention rule lock configuration.
	LockConfiguration *types.LockConfiguration

	// Specifies the resource tags to use to identify resources that are to be
	// retained by a tag-level retention rule. For tag-level retention rules, only
	// deleted resources, of the specified resource type, that have one or more of the
	// specified tag key and value pairs are retained. If a resource is deleted, but it
	// does not have any of the specified tag key and value pairs, it is immediately
	// deleted without being retained by the retention rule.
	//
	// You can add the same tag key and value pair to a maximum or five retention
	// rules.
	//
	// To create a Region-level retention rule, omit this parameter. A Region-level
	// retention rule does not have any resource tags specified. It retains all deleted
	// resources of the specified resource type in the Region in which the rule is
	// created, even if the resources are not tagged.
	ResourceTags []types.ResourceTag

	// Information about the tags to assign to the retention rule.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateRuleOutput struct {

	// The retention rule description.
	Description *string

	// The unique ID of the retention rule.
	Identifier *string

	// Information about the retention rule lock configuration.
	LockConfiguration *types.LockConfiguration

	// The lock state for the retention rule.
	//
	//   - locked - The retention rule is locked and can't be modified or deleted.
	//
	//   - pending_unlock - The retention rule has been unlocked but it is still within
	//   the unlock delay period. The retention rule can be modified or deleted only
	//   after the unlock delay period has expired.
	//
	//   - unlocked - The retention rule is unlocked and it can be modified or deleted
	//   by any user with the required permissions.
	//
	//   - null - The retention rule has never been locked. Once a retention rule has
	//   been locked, it can transition between the locked and unlocked states only; it
	//   can never transition back to null .
	LockState types.LockState

	// Information about the resource tags used to identify resources that are
	// retained by the retention rule.
	ResourceTags []types.ResourceTag

	// The resource type retained by the retention rule.
	ResourceType types.ResourceType

	// Information about the retention period for which the retention rule is to
	// retain resources.
	RetentionPeriod *types.RetentionPeriod

	// The Amazon Resource Name (ARN) of the retention rule.
	RuleArn *string

	// The state of the retention rule. Only retention rules that are in the available
	// state retain resources.
	Status types.RuleStatus

	// Information about the tags assigned to the retention rule.
	Tags []types.Tag

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateRuleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateRule{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateRule{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateRule"); err != nil {
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
	if err = addOpCreateRuleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateRule(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateRule(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateRule",
	}
}

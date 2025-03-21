// Code generated by smithy-go-codegen DO NOT EDIT.

package chimesdkvoice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/chimesdkvoice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Validates an address to be used for 911 calls made with Amazon Chime SDK Voice
// Connectors. You can use validated addresses in a Presence Information Data
// Format Location Object file that you include in SIP requests. That helps ensure
// that addresses are routed to the appropriate Public Safety Answering Point.
func (c *Client) ValidateE911Address(ctx context.Context, params *ValidateE911AddressInput, optFns ...func(*Options)) (*ValidateE911AddressOutput, error) {
	if params == nil {
		params = &ValidateE911AddressInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ValidateE911Address", params, optFns, c.addOperationValidateE911AddressMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ValidateE911AddressOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ValidateE911AddressInput struct {

	// The AWS account ID.
	//
	// This member is required.
	AwsAccountId *string

	// The address city, such as Portland .
	//
	// This member is required.
	City *string

	// The country in the address being validated.
	//
	// This member is required.
	Country *string

	// The dress postal code, such 04352 .
	//
	// This member is required.
	PostalCode *string

	// The address state, such as ME .
	//
	// This member is required.
	State *string

	// The address street information, such as 8th Avenue .
	//
	// This member is required.
	StreetInfo *string

	// The address street number, such as 200 or 2121 .
	//
	// This member is required.
	StreetNumber *string

	noSmithyDocumentSerde
}

type ValidateE911AddressOutput struct {

	// The validated address.
	Address *types.Address

	// The ID that represents the address.
	AddressExternalId *string

	// The list of address suggestions..
	CandidateAddressList []types.CandidateAddress

	// Number indicating the result of address validation. 0 means the address was
	// perfect as-is and successfully validated. 1 means the address was corrected. 2
	// means the address sent was not close enough and was not validated.
	ValidationResult int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationValidateE911AddressMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpValidateE911Address{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpValidateE911Address{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ValidateE911Address"); err != nil {
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
	if err = addOpValidateE911AddressValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opValidateE911Address(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opValidateE911Address(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ValidateE911Address",
	}
}

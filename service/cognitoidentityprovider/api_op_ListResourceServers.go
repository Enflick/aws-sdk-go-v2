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

// Lists the resource servers for a user pool.
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
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [Signing Amazon Web Services API Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html
func (c *Client) ListResourceServers(ctx context.Context, params *ListResourceServersInput, optFns ...func(*Options)) (*ListResourceServersOutput, error) {
	if params == nil {
		params = &ListResourceServersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListResourceServers", params, optFns, c.addOperationListResourceServersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListResourceServersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListResourceServersInput struct {

	// The user pool ID for the user pool.
	//
	// This member is required.
	UserPoolId *string

	// The maximum number of resource servers to return.
	MaxResults *int32

	// A pagination token.
	NextToken *string

	noSmithyDocumentSerde
}

type ListResourceServersOutput struct {

	// The resource servers.
	//
	// This member is required.
	ResourceServers []types.ResourceServerType

	// A pagination token.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListResourceServersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListResourceServers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListResourceServers{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListResourceServers"); err != nil {
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
	if err = addOpListResourceServersValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListResourceServers(options.Region), middleware.Before); err != nil {
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

// ListResourceServersAPIClient is a client that implements the
// ListResourceServers operation.
type ListResourceServersAPIClient interface {
	ListResourceServers(context.Context, *ListResourceServersInput, ...func(*Options)) (*ListResourceServersOutput, error)
}

var _ ListResourceServersAPIClient = (*Client)(nil)

// ListResourceServersPaginatorOptions is the paginator options for
// ListResourceServers
type ListResourceServersPaginatorOptions struct {
	// The maximum number of resource servers to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListResourceServersPaginator is a paginator for ListResourceServers
type ListResourceServersPaginator struct {
	options   ListResourceServersPaginatorOptions
	client    ListResourceServersAPIClient
	params    *ListResourceServersInput
	nextToken *string
	firstPage bool
}

// NewListResourceServersPaginator returns a new ListResourceServersPaginator
func NewListResourceServersPaginator(client ListResourceServersAPIClient, params *ListResourceServersInput, optFns ...func(*ListResourceServersPaginatorOptions)) *ListResourceServersPaginator {
	if params == nil {
		params = &ListResourceServersInput{}
	}

	options := ListResourceServersPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListResourceServersPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListResourceServersPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListResourceServers page.
func (p *ListResourceServersPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListResourceServersOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.ListResourceServers(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListResourceServers(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListResourceServers",
	}
}

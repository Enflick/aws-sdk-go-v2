// Code generated by smithy-go-codegen DO NOT EDIT.

package datasync

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/datasync/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates some of the parameters of a Server Message Block (SMB) file server
// location that you can use for DataSync transfers.
func (c *Client) UpdateLocationSmb(ctx context.Context, params *UpdateLocationSmbInput, optFns ...func(*Options)) (*UpdateLocationSmbOutput, error) {
	if params == nil {
		params = &UpdateLocationSmbInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateLocationSmb", params, optFns, c.addOperationUpdateLocationSmbMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateLocationSmbOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateLocationSmbInput struct {

	// Specifies the ARN of the SMB location that you want to update.
	//
	// This member is required.
	LocationArn *string

	// Specifies the DataSync agent (or agents) which you want to connect to your SMB
	// file server. You specify an agent by using its Amazon Resource Name (ARN).
	AgentArns []string

	// Specifies the Windows domain name that your SMB file server belongs to.
	//
	// If you have multiple domains in your environment, configuring this parameter
	// makes sure that DataSync connects to the right file server.
	//
	// For more information, see [required permissions] for SMB locations.
	//
	// [required permissions]: https://docs.aws.amazon.com/datasync/latest/userguide/create-smb-location.html#configuring-smb-permissions
	Domain *string

	// Specifies the version of the Server Message Block (SMB) protocol that DataSync
	// uses to access an SMB file server.
	MountOptions *types.SmbMountOptions

	// Specifies the password of the user who can mount your SMB file server and has
	// permission to access the files and folders involved in your transfer.
	//
	// For more information, see [required permissions] for SMB locations.
	//
	// [required permissions]: https://docs.aws.amazon.com/datasync/latest/userguide/create-smb-location.html#configuring-smb-permissions
	Password *string

	// Specifies the name of the share exported by your SMB file server where DataSync
	// will read or write data. You can include a subdirectory in the share path (for
	// example, /path/to/subdirectory ). Make sure that other SMB clients in your
	// network can also mount this path.
	//
	// To copy all data in the specified subdirectory, DataSync must be able to mount
	// the SMB share and access all of its data. For more information, see [required permissions]for SMB
	// locations.
	//
	// [required permissions]: https://docs.aws.amazon.com/datasync/latest/userguide/create-smb-location.html#configuring-smb-permissions
	Subdirectory *string

	// Specifies the user name that can mount your SMB file server and has permission
	// to access the files and folders involved in your transfer.
	//
	// For information about choosing a user with the right level of access for your
	// transfer, see [required permissions]for SMB locations.
	//
	// [required permissions]: https://docs.aws.amazon.com/datasync/latest/userguide/create-smb-location.html#configuring-smb-permissions
	User *string

	noSmithyDocumentSerde
}

type UpdateLocationSmbOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateLocationSmbMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateLocationSmb{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateLocationSmb{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateLocationSmb"); err != nil {
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
	if err = addOpUpdateLocationSmbValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateLocationSmb(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateLocationSmb(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateLocationSmb",
	}
}

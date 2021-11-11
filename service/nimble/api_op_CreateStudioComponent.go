// Code generated by smithy-go-codegen DO NOT EDIT.

package nimble

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/nimble/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a studio component resource.
func (c *Client) CreateStudioComponent(ctx context.Context, params *CreateStudioComponentInput, optFns ...func(*Options)) (*CreateStudioComponentOutput, error) {
	if params == nil {
		params = &CreateStudioComponentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateStudioComponent", params, optFns, c.addOperationCreateStudioComponentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateStudioComponentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type CreateStudioComponentInput struct {

	// The name for the studio component.
	//
	// This member is required.
	Name *string

	// The studio ID.
	//
	// This member is required.
	StudioId *string

	// The type of the studio component.
	//
	// This member is required.
	Type types.StudioComponentType

	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request. If you don’t specify a client token, the AWS SDK automatically
	// generates a client token and uses it for the request to ensure idempotency.
	ClientToken *string

	// The configuration of the studio component, based on component type.
	Configuration *types.StudioComponentConfiguration

	// The description.
	Description *string

	// The EC2 security groups that control access to the studio component.
	Ec2SecurityGroupIds []string

	// Initialization scripts for studio components.
	InitializationScripts []types.StudioComponentInitializationScript

	// Parameters for the studio component scripts.
	ScriptParameters []types.ScriptParameterKeyValue

	// The specific subtype of a studio component.
	Subtype types.StudioComponentSubtype

	// A collection of labels, in the form of key:value pairs, that apply to this
	// resource.
	Tags map[string]string

	noSmithyDocumentSerde
}

//
type CreateStudioComponentOutput struct {

	// Information about the studio component.
	StudioComponent *types.StudioComponent

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateStudioComponentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateStudioComponent{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateStudioComponent{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpCreateStudioComponentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateStudioComponent(options.Region), middleware.Before); err != nil {
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
	return nil
}

func newServiceMetadataMiddleware_opCreateStudioComponent(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "nimble",
		OperationName: "CreateStudioComponent",
	}
}
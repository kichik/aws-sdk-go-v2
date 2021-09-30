// Code generated by smithy-go-codegen DO NOT EDIT.

package proton

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/proton/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Attempts to cancel a service instance deployment on an UpdateServiceInstance
// action, if the deployment is IN_PROGRESS. For more information, see Update a
// service instance in the AWS Proton Administrator guide
// (https://docs.aws.amazon.com/proton/latest/adminguide/ag-svc-instance-update.html)
// or the AWS Proton User guide
// (https://docs.aws.amazon.com/proton/latest/userguide/ug-svc-instance-update.html).
// The following list includes potential cancellation scenarios.
//
// * If the
// cancellation attempt succeeds, the resulting deployment state is CANCELLED.
//
// *
// If the cancellation attempt fails, the resulting deployment state is FAILED.
//
// *
// If the current UpdateServiceInstance action succeeds before the cancellation
// attempt starts, the resulting deployment state is SUCCEEDED and the cancellation
// attempt has no effect.
func (c *Client) CancelServiceInstanceDeployment(ctx context.Context, params *CancelServiceInstanceDeploymentInput, optFns ...func(*Options)) (*CancelServiceInstanceDeploymentOutput, error) {
	if params == nil {
		params = &CancelServiceInstanceDeploymentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CancelServiceInstanceDeployment", params, optFns, c.addOperationCancelServiceInstanceDeploymentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CancelServiceInstanceDeploymentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CancelServiceInstanceDeploymentInput struct {

	// The name of the service instance with the deployment to cancel.
	//
	// This member is required.
	ServiceInstanceName *string

	// The name of the service with the service instance deployment to cancel.
	//
	// This member is required.
	ServiceName *string

	noSmithyDocumentSerde
}

type CancelServiceInstanceDeploymentOutput struct {

	// The service instance summary data that's returned by AWS Proton.
	//
	// This member is required.
	ServiceInstance *types.ServiceInstance

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCancelServiceInstanceDeploymentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpCancelServiceInstanceDeployment{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpCancelServiceInstanceDeployment{}, middleware.After)
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
	if err = addOpCancelServiceInstanceDeploymentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCancelServiceInstanceDeployment(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCancelServiceInstanceDeployment(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "proton",
		OperationName: "CancelServiceInstanceDeployment",
	}
}
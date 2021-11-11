// Code generated by smithy-go-codegen DO NOT EDIT.

package location

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/location/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieves the map resource details.
func (c *Client) DescribeMap(ctx context.Context, params *DescribeMapInput, optFns ...func(*Options)) (*DescribeMapOutput, error) {
	if params == nil {
		params = &DescribeMapInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeMap", params, optFns, c.addOperationDescribeMapMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeMapOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeMapInput struct {

	// The name of the map resource.
	//
	// This member is required.
	MapName *string

	noSmithyDocumentSerde
}

type DescribeMapOutput struct {

	// Specifies the map tile style selected from a partner data provider.
	//
	// This member is required.
	Configuration *types.MapConfiguration

	// The timestamp for when the map resource was created in ISO 8601
	// (https://www.iso.org/iso-8601-date-and-time-format.html) format:
	// YYYY-MM-DDThh:mm:ss.sssZ.
	//
	// This member is required.
	CreateTime *time.Time

	// Specifies the data provider for the associated map tiles.
	//
	// This member is required.
	DataSource *string

	// The optional description for the map resource.
	//
	// This member is required.
	Description *string

	// The Amazon Resource Name (ARN) for the map resource. Used to specify a resource
	// across all AWS.
	//
	// * Format example: arn:aws:geo:region:account-id:maps/ExampleMap
	//
	// This member is required.
	MapArn *string

	// The map style selected from an available provider.
	//
	// This member is required.
	MapName *string

	// The pricing plan selected for the specified map resource. For additional details
	// and restrictions on each pricing plan option, see Amazon Location Service
	// pricing (https://aws.amazon.com/location/pricing/).
	//
	// This member is required.
	PricingPlan types.PricingPlan

	// The timestamp for when the map resource was last update in ISO 8601
	// (https://www.iso.org/iso-8601-date-and-time-format.html) format:
	// YYYY-MM-DDThh:mm:ss.sssZ.
	//
	// This member is required.
	UpdateTime *time.Time

	// Tags associated with the map resource.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeMapMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeMap{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeMap{}, middleware.After)
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
	if err = addRetryMiddlewares(stack, options); err != nil {
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
	if err = addOpDescribeMapValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeMap(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeMap(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeMap",
	}
}
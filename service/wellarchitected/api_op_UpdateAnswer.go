// Code generated by smithy-go-codegen DO NOT EDIT.

package wellarchitected

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/wellarchitected/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Update the answer to a specific question in a workload review.
func (c *Client) UpdateAnswer(ctx context.Context, params *UpdateAnswerInput, optFns ...func(*Options)) (*UpdateAnswerOutput, error) {
	if params == nil {
		params = &UpdateAnswerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateAnswer", params, optFns, c.addOperationUpdateAnswerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateAnswerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Input to update answer.
type UpdateAnswerInput struct {

	// The alias of the lens, for example, serverless. Each lens is identified by its
	// LensSummary$LensAlias.
	//
	// This member is required.
	LensAlias *string

	// The ID of the question.
	//
	// This member is required.
	QuestionId *string

	// The ID assigned to the workload. This ID is unique within an AWS Region.
	//
	// This member is required.
	WorkloadId *string

	// A list of choices to update on a question in your workload. The String key
	// corresponds to the choice ID to be updated.
	ChoiceUpdates map[string]types.ChoiceUpdate

	// Defines whether this question is applicable to a lens review.
	IsApplicable bool

	// The notes associated with the workload.
	Notes *string

	// The reason why a question is not applicable to your workload.
	Reason types.AnswerReason

	// List of selected choice IDs in a question answer. The values entered replace the
	// previously selected choices.
	SelectedChoices []string

	noSmithyDocumentSerde
}

// Output of a update answer call.
type UpdateAnswerOutput struct {

	// An answer of the question.
	Answer *types.Answer

	// The alias of the lens, for example, serverless. Each lens is identified by its
	// LensSummary$LensAlias.
	LensAlias *string

	// The ID assigned to the workload. This ID is unique within an AWS Region.
	WorkloadId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateAnswerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateAnswer{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateAnswer{}, middleware.After)
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
	if err = addOpUpdateAnswerValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateAnswer(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateAnswer(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "wellarchitected",
		OperationName: "UpdateAnswer",
	}
}
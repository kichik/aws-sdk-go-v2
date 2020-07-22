// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package awsendpointdiscoverytest

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type TestDiscoveryIdentifiersRequiredInput struct {
	_ struct{} `type:"structure"`

	// Sdk is a required field
	Sdk *string `type:"string" required:"true"`
}

// String returns the string representation
func (s TestDiscoveryIdentifiersRequiredInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *TestDiscoveryIdentifiersRequiredInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "TestDiscoveryIdentifiersRequiredInput"}

	if s.Sdk == nil {
		invalidParams.Add(aws.NewErrParamRequired("Sdk"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type TestDiscoveryIdentifiersRequiredOutput struct {
	_ struct{} `type:"structure"`

	RequestSuccessful *bool `type:"boolean"`
}

// String returns the string representation
func (s TestDiscoveryIdentifiersRequiredOutput) String() string {
	return awsutil.Prettify(s)
}

const opTestDiscoveryIdentifiersRequired = "TestDiscoveryIdentifiersRequired"

// TestDiscoveryIdentifiersRequiredRequest returns a request value for making API operation for
// AwsEndpointDiscoveryTest.
//
//    // Example sending a request using TestDiscoveryIdentifiersRequiredRequest.
//    req := client.TestDiscoveryIdentifiersRequiredRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) TestDiscoveryIdentifiersRequiredRequest(input *TestDiscoveryIdentifiersRequiredInput) TestDiscoveryIdentifiersRequiredRequest {
	op := &aws.Operation{
		Name:       opTestDiscoveryIdentifiersRequired,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &TestDiscoveryIdentifiersRequiredInput{}
	}

	req := c.newRequest(op, input, &TestDiscoveryIdentifiersRequiredOutput{})

	de := discovererDescribeEndpoints{
		Client:        c,
		Required:      true,
		EndpointCache: c.endpointCache,
		Params: map[string]*string{
			"op":  &req.Operation.Name,
			"Sdk": input.Sdk,
		},
	}

	for k, v := range de.Params {
		if v == nil {
			delete(de.Params, k)
		}
	}

	req.Handlers.Build.PushFrontNamed(aws.NamedHandler{
		Name: "crr.endpointdiscovery",
		Fn:   de.Handler,
	})

	return TestDiscoveryIdentifiersRequiredRequest{Request: req, Input: input, Copy: c.TestDiscoveryIdentifiersRequiredRequest}
}

// TestDiscoveryIdentifiersRequiredRequest is the request type for the
// TestDiscoveryIdentifiersRequired API operation.
type TestDiscoveryIdentifiersRequiredRequest struct {
	*aws.Request
	Input *TestDiscoveryIdentifiersRequiredInput
	Copy  func(*TestDiscoveryIdentifiersRequiredInput) TestDiscoveryIdentifiersRequiredRequest
}

// Send marshals and sends the TestDiscoveryIdentifiersRequired API request.
func (r TestDiscoveryIdentifiersRequiredRequest) Send(ctx context.Context) (*TestDiscoveryIdentifiersRequiredResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &TestDiscoveryIdentifiersRequiredResponse{
		TestDiscoveryIdentifiersRequiredOutput: r.Request.Data.(*TestDiscoveryIdentifiersRequiredOutput),
		response:                               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// TestDiscoveryIdentifiersRequiredResponse is the response type for the
// TestDiscoveryIdentifiersRequired API operation.
type TestDiscoveryIdentifiersRequiredResponse struct {
	*TestDiscoveryIdentifiersRequiredOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// TestDiscoveryIdentifiersRequired request.
func (r *TestDiscoveryIdentifiersRequiredResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
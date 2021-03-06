// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sqs

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteMessageBatchInput struct {
	_ struct{} `type:"structure"`

	// A list of receipt handles for the messages to be deleted.
	//
	// Entries is a required field
	Entries []DeleteMessageBatchRequestEntry `locationNameList:"DeleteMessageBatchRequestEntry" type:"list" flattened:"true" required:"true"`

	// The URL of the Amazon SQS queue from which messages are deleted.
	//
	// Queue URLs and names are case-sensitive.
	//
	// QueueUrl is a required field
	QueueUrl *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteMessageBatchInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteMessageBatchInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteMessageBatchInput"}

	if s.Entries == nil {
		invalidParams.Add(aws.NewErrParamRequired("Entries"))
	}

	if s.QueueUrl == nil {
		invalidParams.Add(aws.NewErrParamRequired("QueueUrl"))
	}
	if s.Entries != nil {
		for i, v := range s.Entries {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Entries", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// For each message in the batch, the response contains a DeleteMessageBatchResultEntry
// tag if the message is deleted or a BatchResultErrorEntry tag if the message
// can't be deleted.
type DeleteMessageBatchOutput struct {
	_ struct{} `type:"structure"`

	// A list of BatchResultErrorEntry items.
	//
	// Failed is a required field
	Failed []BatchResultErrorEntry `locationNameList:"BatchResultErrorEntry" type:"list" flattened:"true" required:"true"`

	// A list of DeleteMessageBatchResultEntry items.
	//
	// Successful is a required field
	Successful []DeleteMessageBatchResultEntry `locationNameList:"DeleteMessageBatchResultEntry" type:"list" flattened:"true" required:"true"`
}

// String returns the string representation
func (s DeleteMessageBatchOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteMessageBatch = "DeleteMessageBatch"

// DeleteMessageBatchRequest returns a request value for making API operation for
// Amazon Simple Queue Service.
//
// Deletes up to ten messages from the specified queue. This is a batch version
// of DeleteMessage. The result of the action on each message is reported individually
// in the response.
//
// Because the batch request can result in a combination of successful and unsuccessful
// actions, you should check for batch errors even when the call returns an
// HTTP status code of 200.
//
// Some actions take lists of parameters. These lists are specified using the
// param.n notation. Values of n are integers starting from 1. For example,
// a parameter list with two elements looks like this:
//
// &Attribute.1=first
//
// &Attribute.2=second
//
//    // Example sending a request using DeleteMessageBatchRequest.
//    req := client.DeleteMessageBatchRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sqs-2012-11-05/DeleteMessageBatch
func (c *Client) DeleteMessageBatchRequest(input *DeleteMessageBatchInput) DeleteMessageBatchRequest {
	op := &aws.Operation{
		Name:       opDeleteMessageBatch,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteMessageBatchInput{}
	}

	req := c.newRequest(op, input, &DeleteMessageBatchOutput{})

	return DeleteMessageBatchRequest{Request: req, Input: input, Copy: c.DeleteMessageBatchRequest}
}

// DeleteMessageBatchRequest is the request type for the
// DeleteMessageBatch API operation.
type DeleteMessageBatchRequest struct {
	*aws.Request
	Input *DeleteMessageBatchInput
	Copy  func(*DeleteMessageBatchInput) DeleteMessageBatchRequest
}

// Send marshals and sends the DeleteMessageBatch API request.
func (r DeleteMessageBatchRequest) Send(ctx context.Context) (*DeleteMessageBatchResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteMessageBatchResponse{
		DeleteMessageBatchOutput: r.Request.Data.(*DeleteMessageBatchOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteMessageBatchResponse is the response type for the
// DeleteMessageBatch API operation.
type DeleteMessageBatchResponse struct {
	*DeleteMessageBatchOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteMessageBatch request.
func (r *DeleteMessageBatchResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

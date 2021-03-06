// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package directoryservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Contains the inputs for the DescribeSnapshots operation.
type DescribeSnapshotsInput struct {
	_ struct{} `type:"structure"`

	// The identifier of the directory for which to retrieve snapshot information.
	DirectoryId *string `type:"string"`

	// The maximum number of objects to return.
	Limit *int64 `type:"integer"`

	// The DescribeSnapshotsResult.NextToken value from a previous call to DescribeSnapshots.
	// Pass null if this is the first call.
	NextToken *string `type:"string"`

	// A list of identifiers of the snapshots to obtain the information for. If
	// this member is null or empty, all snapshots are returned using the Limit
	// and NextToken members.
	SnapshotIds []string `type:"list"`
}

// String returns the string representation
func (s DescribeSnapshotsInput) String() string {
	return awsutil.Prettify(s)
}

// Contains the results of the DescribeSnapshots operation.
type DescribeSnapshotsOutput struct {
	_ struct{} `type:"structure"`

	// If not null, more results are available. Pass this value in the NextToken
	// member of a subsequent call to DescribeSnapshots.
	NextToken *string `type:"string"`

	// The list of Snapshot objects that were retrieved.
	//
	// It is possible that this list contains less than the number of items specified
	// in the Limit member of the request. This occurs if there are less than the
	// requested number of items left to retrieve, or if the limitations of the
	// operation have been exceeded.
	Snapshots []Snapshot `type:"list"`
}

// String returns the string representation
func (s DescribeSnapshotsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeSnapshots = "DescribeSnapshots"

// DescribeSnapshotsRequest returns a request value for making API operation for
// AWS Directory Service.
//
// Obtains information about the directory snapshots that belong to this account.
//
// This operation supports pagination with the use of the NextToken request
// and response parameters. If more results are available, the DescribeSnapshots.NextToken
// member contains a token that you pass in the next call to DescribeSnapshots
// to retrieve the next set of items.
//
// You can also specify a maximum number of return results with the Limit parameter.
//
//    // Example sending a request using DescribeSnapshotsRequest.
//    req := client.DescribeSnapshotsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ds-2015-04-16/DescribeSnapshots
func (c *Client) DescribeSnapshotsRequest(input *DescribeSnapshotsInput) DescribeSnapshotsRequest {
	op := &aws.Operation{
		Name:       opDescribeSnapshots,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeSnapshotsInput{}
	}

	req := c.newRequest(op, input, &DescribeSnapshotsOutput{})

	return DescribeSnapshotsRequest{Request: req, Input: input, Copy: c.DescribeSnapshotsRequest}
}

// DescribeSnapshotsRequest is the request type for the
// DescribeSnapshots API operation.
type DescribeSnapshotsRequest struct {
	*aws.Request
	Input *DescribeSnapshotsInput
	Copy  func(*DescribeSnapshotsInput) DescribeSnapshotsRequest
}

// Send marshals and sends the DescribeSnapshots API request.
func (r DescribeSnapshotsRequest) Send(ctx context.Context) (*DescribeSnapshotsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeSnapshotsResponse{
		DescribeSnapshotsOutput: r.Request.Data.(*DescribeSnapshotsOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeSnapshotsResponse is the response type for the
// DescribeSnapshots API operation.
type DescribeSnapshotsResponse struct {
	*DescribeSnapshotsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeSnapshots request.
func (r *DescribeSnapshotsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

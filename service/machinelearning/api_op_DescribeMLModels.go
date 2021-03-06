// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package machinelearning

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeMLModelsInput struct {
	_ struct{} `type:"structure"`

	// The equal to operator. The MLModel results will have FilterVariable values
	// that exactly match the value specified with EQ.
	EQ *string `type:"string"`

	// Use one of the following variables to filter a list of MLModel:
	//
	//    * CreatedAt - Sets the search criteria to MLModel creation date.
	//
	//    * Status - Sets the search criteria to MLModel status.
	//
	//    * Name - Sets the search criteria to the contents of MLModel Name.
	//
	//    * IAMUser - Sets the search criteria to the user account that invoked
	//    the MLModel creation.
	//
	//    * TrainingDataSourceId - Sets the search criteria to the DataSource used
	//    to train one or more MLModel.
	//
	//    * RealtimeEndpointStatus - Sets the search criteria to the MLModel real-time
	//    endpoint status.
	//
	//    * MLModelType - Sets the search criteria to MLModel type: binary, regression,
	//    or multi-class.
	//
	//    * Algorithm - Sets the search criteria to the algorithm that the MLModel
	//    uses.
	//
	//    * TrainingDataURI - Sets the search criteria to the data file(s) used
	//    in training a MLModel. The URL can identify either a file or an Amazon
	//    Simple Storage Service (Amazon S3) bucket or directory.
	FilterVariable MLModelFilterVariable `type:"string" enum:"true"`

	// The greater than or equal to operator. The MLModel results will have FilterVariable
	// values that are greater than or equal to the value specified with GE.
	GE *string `type:"string"`

	// The greater than operator. The MLModel results will have FilterVariable values
	// that are greater than the value specified with GT.
	GT *string `type:"string"`

	// The less than or equal to operator. The MLModel results will have FilterVariable
	// values that are less than or equal to the value specified with LE.
	LE *string `type:"string"`

	// The less than operator. The MLModel results will have FilterVariable values
	// that are less than the value specified with LT.
	LT *string `type:"string"`

	// The number of pages of information to include in the result. The range of
	// acceptable values is 1 through 100. The default value is 100.
	Limit *int64 `min:"1" type:"integer"`

	// The not equal to operator. The MLModel results will have FilterVariable values
	// not equal to the value specified with NE.
	NE *string `type:"string"`

	// The ID of the page in the paginated results.
	NextToken *string `type:"string"`

	// A string that is found at the beginning of a variable, such as Name or Id.
	//
	// For example, an MLModel could have the Name 2014-09-09-HolidayGiftMailer.
	// To search for this MLModel, select Name for the FilterVariable and any of
	// the following strings for the Prefix:
	//
	//    * 2014-09
	//
	//    * 2014-09-09
	//
	//    * 2014-09-09-Holiday
	Prefix *string `type:"string"`

	// A two-value parameter that determines the sequence of the resulting list
	// of MLModel.
	//
	//    * asc - Arranges the list in ascending order (A-Z, 0-9).
	//
	//    * dsc - Arranges the list in descending order (Z-A, 9-0).
	//
	// Results are sorted by FilterVariable.
	SortOrder SortOrder `type:"string" enum:"true"`
}

// String returns the string representation
func (s DescribeMLModelsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeMLModelsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeMLModelsInput"}
	if s.Limit != nil && *s.Limit < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("Limit", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the output of a DescribeMLModels operation. The content is essentially
// a list of MLModel.
type DescribeMLModelsOutput struct {
	_ struct{} `type:"structure"`

	// The ID of the next page in the paginated results that indicates at least
	// one more page follows.
	NextToken *string `type:"string"`

	// A list of MLModel that meet the search criteria.
	Results []MLModel `type:"list"`
}

// String returns the string representation
func (s DescribeMLModelsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeMLModels = "DescribeMLModels"

// DescribeMLModelsRequest returns a request value for making API operation for
// Amazon Machine Learning.
//
// Returns a list of MLModel that match the search criteria in the request.
//
//    // Example sending a request using DescribeMLModelsRequest.
//    req := client.DescribeMLModelsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) DescribeMLModelsRequest(input *DescribeMLModelsInput) DescribeMLModelsRequest {
	op := &aws.Operation{
		Name:       opDescribeMLModels,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "Limit",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeMLModelsInput{}
	}

	req := c.newRequest(op, input, &DescribeMLModelsOutput{})

	return DescribeMLModelsRequest{Request: req, Input: input, Copy: c.DescribeMLModelsRequest}
}

// DescribeMLModelsRequest is the request type for the
// DescribeMLModels API operation.
type DescribeMLModelsRequest struct {
	*aws.Request
	Input *DescribeMLModelsInput
	Copy  func(*DescribeMLModelsInput) DescribeMLModelsRequest
}

// Send marshals and sends the DescribeMLModels API request.
func (r DescribeMLModelsRequest) Send(ctx context.Context) (*DescribeMLModelsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeMLModelsResponse{
		DescribeMLModelsOutput: r.Request.Data.(*DescribeMLModelsOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeMLModelsRequestPaginator returns a paginator for DescribeMLModels.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeMLModelsRequest(input)
//   p := machinelearning.NewDescribeMLModelsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeMLModelsPaginator(req DescribeMLModelsRequest) DescribeMLModelsPaginator {
	return DescribeMLModelsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeMLModelsInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// DescribeMLModelsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeMLModelsPaginator struct {
	aws.Pager
}

func (p *DescribeMLModelsPaginator) CurrentPage() *DescribeMLModelsOutput {
	return p.Pager.CurrentPage().(*DescribeMLModelsOutput)
}

// DescribeMLModelsResponse is the response type for the
// DescribeMLModels API operation.
type DescribeMLModelsResponse struct {
	*DescribeMLModelsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeMLModels request.
func (r *DescribeMLModelsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

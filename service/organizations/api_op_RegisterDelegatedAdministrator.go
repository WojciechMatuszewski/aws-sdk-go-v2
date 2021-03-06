// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package organizations

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

type RegisterDelegatedAdministratorInput struct {
	_ struct{} `type:"structure"`

	// The account ID number of the member account in the organization to register
	// as a delegated administrator.
	//
	// AccountId is a required field
	AccountId *string `type:"string" required:"true"`

	// The service principal of the AWS service for which you want to make the member
	// account a delegated administrator.
	//
	// ServicePrincipal is a required field
	ServicePrincipal *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s RegisterDelegatedAdministratorInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RegisterDelegatedAdministratorInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RegisterDelegatedAdministratorInput"}

	if s.AccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccountId"))
	}

	if s.ServicePrincipal == nil {
		invalidParams.Add(aws.NewErrParamRequired("ServicePrincipal"))
	}
	if s.ServicePrincipal != nil && len(*s.ServicePrincipal) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ServicePrincipal", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type RegisterDelegatedAdministratorOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s RegisterDelegatedAdministratorOutput) String() string {
	return awsutil.Prettify(s)
}

const opRegisterDelegatedAdministrator = "RegisterDelegatedAdministrator"

// RegisterDelegatedAdministratorRequest returns a request value for making API operation for
// AWS Organizations.
//
// Enables the specified member account to administer the Organizations features
// of the specified AWS service. It grants read-only access to AWS Organizations
// service data. The account still requires IAM permissions to access and administer
// the AWS service.
//
// You can run this action only for AWS services that support this feature.
// For a current list of services that support it, see the column Supports Delegated
// Administrator in the table at AWS Services that you can use with AWS Organizations
// (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_integrated-services-list.html)
// in the AWS Organizations User Guide.
//
// This operation can be called only from the organization's master account.
//
//    // Example sending a request using RegisterDelegatedAdministratorRequest.
//    req := client.RegisterDelegatedAdministratorRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/organizations-2016-11-28/RegisterDelegatedAdministrator
func (c *Client) RegisterDelegatedAdministratorRequest(input *RegisterDelegatedAdministratorInput) RegisterDelegatedAdministratorRequest {
	op := &aws.Operation{
		Name:       opRegisterDelegatedAdministrator,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RegisterDelegatedAdministratorInput{}
	}

	req := c.newRequest(op, input, &RegisterDelegatedAdministratorOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return RegisterDelegatedAdministratorRequest{Request: req, Input: input, Copy: c.RegisterDelegatedAdministratorRequest}
}

// RegisterDelegatedAdministratorRequest is the request type for the
// RegisterDelegatedAdministrator API operation.
type RegisterDelegatedAdministratorRequest struct {
	*aws.Request
	Input *RegisterDelegatedAdministratorInput
	Copy  func(*RegisterDelegatedAdministratorInput) RegisterDelegatedAdministratorRequest
}

// Send marshals and sends the RegisterDelegatedAdministrator API request.
func (r RegisterDelegatedAdministratorRequest) Send(ctx context.Context) (*RegisterDelegatedAdministratorResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RegisterDelegatedAdministratorResponse{
		RegisterDelegatedAdministratorOutput: r.Request.Data.(*RegisterDelegatedAdministratorOutput),
		response:                             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RegisterDelegatedAdministratorResponse is the response type for the
// RegisterDelegatedAdministrator API operation.
type RegisterDelegatedAdministratorResponse struct {
	*RegisterDelegatedAdministratorOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RegisterDelegatedAdministrator request.
func (r *RegisterDelegatedAdministratorResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package fsx

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// The request object for the CreateFileSystemFromBackup operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/fsx-2018-03-01/CreateFileSystemFromBackupRequest
type CreateFileSystemFromBackupInput struct {
	_ struct{} `type:"structure"`

	// The ID of the backup.
	//
	// BackupId is a required field
	BackupId *string `min:"12" type:"string" required:"true"`

	// (Optional) A string of up to 64 ASCII characters that Amazon FSx uses to
	// ensure idempotent creation. This string is automatically filled on your behalf
	// when you use the AWS Command Line Interface (AWS CLI) or an AWS SDK.
	ClientRequestToken *string `min:"1" type:"string" idempotencyToken:"true"`

	// A list of IDs for the security groups that apply to the specified network
	// interfaces created for file system access. These security groups apply to
	// all network interfaces. This value isn't returned in later describe requests.
	SecurityGroupIds []string `type:"list"`

	// A list of IDs for the subnets that the file system will be accessible from.
	// Currently, you can specify only one subnet. The file server is also launched
	// in that subnet's Availability Zone.
	//
	// SubnetIds is a required field
	SubnetIds []string `type:"list" required:"true"`

	// The tags to be applied to the file system at file system creation. The key
	// value of the Name tag appears in the console as the file system name.
	Tags []Tag `min:"1" type:"list"`

	// The configuration for this Microsoft Windows file system.
	WindowsConfiguration *CreateFileSystemWindowsConfiguration `type:"structure"`
}

// String returns the string representation
func (s CreateFileSystemFromBackupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateFileSystemFromBackupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateFileSystemFromBackupInput"}

	if s.BackupId == nil {
		invalidParams.Add(aws.NewErrParamRequired("BackupId"))
	}
	if s.BackupId != nil && len(*s.BackupId) < 12 {
		invalidParams.Add(aws.NewErrParamMinLen("BackupId", 12))
	}
	if s.ClientRequestToken != nil && len(*s.ClientRequestToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ClientRequestToken", 1))
	}

	if s.SubnetIds == nil {
		invalidParams.Add(aws.NewErrParamRequired("SubnetIds"))
	}
	if s.Tags != nil && len(s.Tags) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Tags", 1))
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.WindowsConfiguration != nil {
		if err := s.WindowsConfiguration.Validate(); err != nil {
			invalidParams.AddNested("WindowsConfiguration", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The response object for the CreateFileSystemFromBackup operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/fsx-2018-03-01/CreateFileSystemFromBackupResponse
type CreateFileSystemFromBackupOutput struct {
	_ struct{} `type:"structure"`

	// A description of the file system.
	FileSystem *FileSystem `type:"structure"`
}

// String returns the string representation
func (s CreateFileSystemFromBackupOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateFileSystemFromBackup = "CreateFileSystemFromBackup"

// CreateFileSystemFromBackupRequest returns a request value for making API operation for
// Amazon FSx.
//
// Creates a new Amazon FSx file system from an existing Amazon FSx for Windows
// File Server backup.
//
// If a file system with the specified client request token exists and the parameters
// match, this call returns the description of the existing file system. If
// a client request token specified by the file system exists and the parameters
// don't match, this call returns IncompatibleParameterError. If a file system
// with the specified client request token doesn't exist, this operation does
// the following:
//
//    * Creates a new Amazon FSx file system from backup with an assigned ID,
//    and an initial lifecycle state of CREATING.
//
//    * Returns the description of the file system.
//
// Parameters like Active Directory, default share name, automatic backup, and
// backup settings default to the parameters of the file system that was backed
// up, unless overridden. You can explicitly supply other settings.
//
// By using the idempotent operation, you can retry a CreateFileSystemFromBackup
// call without the risk of creating an extra file system. This approach can
// be useful when an initial call fails in a way that makes it unclear whether
// a file system was created. Examples are if a transport level timeout occurred,
// or your connection was reset. If you use the same client request token and
// the initial call created a file system, the client receives success as long
// as the parameters are the same.
//
// The CreateFileSystemFromBackup call returns while the file system's lifecycle
// state is still CREATING. You can check the file-system creation status by
// calling the DescribeFileSystems operation, which returns the file system
// state along with other information.
//
//    // Example sending a request using CreateFileSystemFromBackupRequest.
//    req := client.CreateFileSystemFromBackupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/fsx-2018-03-01/CreateFileSystemFromBackup
func (c *Client) CreateFileSystemFromBackupRequest(input *CreateFileSystemFromBackupInput) CreateFileSystemFromBackupRequest {
	op := &aws.Operation{
		Name:       opCreateFileSystemFromBackup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateFileSystemFromBackupInput{}
	}

	req := c.newRequest(op, input, &CreateFileSystemFromBackupOutput{})
	return CreateFileSystemFromBackupRequest{Request: req, Input: input, Copy: c.CreateFileSystemFromBackupRequest}
}

// CreateFileSystemFromBackupRequest is the request type for the
// CreateFileSystemFromBackup API operation.
type CreateFileSystemFromBackupRequest struct {
	*aws.Request
	Input *CreateFileSystemFromBackupInput
	Copy  func(*CreateFileSystemFromBackupInput) CreateFileSystemFromBackupRequest
}

// Send marshals and sends the CreateFileSystemFromBackup API request.
func (r CreateFileSystemFromBackupRequest) Send(ctx context.Context) (*CreateFileSystemFromBackupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateFileSystemFromBackupResponse{
		CreateFileSystemFromBackupOutput: r.Request.Data.(*CreateFileSystemFromBackupOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateFileSystemFromBackupResponse is the response type for the
// CreateFileSystemFromBackup API operation.
type CreateFileSystemFromBackupResponse struct {
	*CreateFileSystemFromBackupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateFileSystemFromBackup request.
func (r *CreateFileSystemFromBackupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
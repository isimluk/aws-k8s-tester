// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Adds or updates an inline policy document that is embedded in the specified IAM
// group. A user can also have managed policies attached to it. To attach a managed
// policy to a group, use AttachGroupPolicy. To create a new managed policy, use
// CreatePolicy. For information about policies, see Managed Policies and Inline
// Policies
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html)
// in the IAM User Guide. For information about limits on the number of inline
// policies that you can embed in a group, see Limitations on IAM Entities
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/LimitationsOnEntities.html) in
// the IAM User Guide. Because policy documents can be large, you should use POST
// rather than GET when calling PutGroupPolicy. For general information about using
// the Query API with IAM, go to Making Query Requests
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/IAM_UsingQueryAPI.html) in the
// IAM User Guide.
func (c *Client) PutGroupPolicy(ctx context.Context, params *PutGroupPolicyInput, optFns ...func(*Options)) (*PutGroupPolicyOutput, error) {
	if params == nil {
		params = &PutGroupPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutGroupPolicy", params, optFns, addOperationPutGroupPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutGroupPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutGroupPolicyInput struct {

	// The name of the group to associate the policy with. This parameter allows
	// (through its regex pattern (http://wikipedia.org/wiki/regex)) a string of
	// characters consisting of upper and lowercase alphanumeric characters with no
	// spaces. You can also include any of the following characters: _+=,.@-.
	//
	// This member is required.
	GroupName *string

	// The policy document. You must provide policies in JSON format in IAM. However,
	// for AWS CloudFormation templates formatted in YAML, you can provide the policy
	// in JSON or YAML format. AWS CloudFormation always converts a YAML policy to JSON
	// format before submitting it to IAM. The regex pattern
	// (http://wikipedia.org/wiki/regex) used to validate this parameter is a string of
	// characters consisting of the following:
	//
	// * Any printable ASCII character ranging
	// from the space character (\u0020) through the end of the ASCII character
	// range
	//
	// * The printable characters in the Basic Latin and Latin-1 Supplement
	// character set (through \u00FF)
	//
	// * The special characters tab (\u0009), line feed
	// (\u000A), and carriage return (\u000D)
	//
	// This member is required.
	PolicyDocument *string

	// The name of the policy document. This parameter allows (through its regex
	// pattern (http://wikipedia.org/wiki/regex)) a string of characters consisting of
	// upper and lowercase alphanumeric characters with no spaces. You can also include
	// any of the following characters: _+=,.@-
	//
	// This member is required.
	PolicyName *string
}

type PutGroupPolicyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPutGroupPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpPutGroupPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpPutGroupPolicy{}, middleware.After)
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
	if err = addOpPutGroupPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutGroupPolicy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutGroupPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "PutGroupPolicy",
	}
}
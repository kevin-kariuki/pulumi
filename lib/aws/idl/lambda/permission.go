// Copyright 2017 Pulumi, Inc. All rights reserved.

package lambda

import (
	"github.com/pulumi/coconut/pkg/resource/idl"

	aws "github.com/pulumi/coconut/lib/aws/idl"
)

// The Permission resource associates a policy statement with a specific AWS Lambda function's access policy.  The
// function policy grants a specific AWS service or application permission to invoke the function.  For more
// information, see http://docs.aws.amazon.com/lambda/latest/dg/API_AddPermission.html.
type Permission struct {
	idl.NamedResource
	// The Lambda actions that you want to allow in this statement.  For example, you can specify lambda:CreateFunction
	// to specify a certain action, or use a wildcard (lambda:*) to grant permission to all Lambda actions.  For a list
	// of actions, see http://docs.aws.amazon.com/IAM/latest/UserGuide/list_lambda.html.
	Action string `coco:"action,replaces"`
	// The Lambda function that you want to associate with this statement.
	Function *Function `coco:"function,replaces"`
	// The entity for which you are granting permission to invoke the Lambda function.  This entity can be any valid AWS
	// service principal, such as `s3.amazonaws.com` or `sns.amazonaws.com`, or, if you are granting cross-account
	// permission, an AWS account ID.  For example, you might want to allow a custom application in another AWS account
	// to push events to Lambda by invoking your function.
	Principal string `coco:"principal,replaces"`
	// The AWS account ID (without hyphens) of the source owner.  For example, if you specify an S3 bucket in the
	// sourceARN property, this value is the bucket owner's account ID.  You can use this property to ensure that all
	// source principals are owned by a specific account.
	SourceAccount *string `coco:"sourceAccount,replaces,optional"`
	// The ARN of a resource that is invoking your function.  When granting Amazon Simple Storage Service (Amazon S3)
	// permission to invoke your function, specify this property with the bucket ARN as its value.  This ensures that
	// events generated only from the specified bucket, not just any bucket from any AWS account that creates a mapping
	// to your function, can invoke the function.
	SourceARN *aws.ARN `coco:"sourceARN,replaces,optional"`
}
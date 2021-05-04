// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sagemaker

import (
	"github.com/aws/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeConflictException for service response error code
	// "ConflictException".
	//
	// There was a conflict when you attempted to modify a SageMaker entity such
	// as an Experiment or Artifact.
	ErrCodeConflictException = "ConflictException"

	// ErrCodeResourceInUse for service response error code
	// "ResourceInUse".
	//
	// The resource being accessed is in use.
	ErrCodeResourceInUse = "ResourceInUse"

	// ErrCodeResourceLimitExceeded for service response error code
	// "ResourceLimitExceeded".
	//
	// You have exceeded an Amazon SageMaker resource limit. For example, you might
	// have too many training jobs created.
	ErrCodeResourceLimitExceeded = "ResourceLimitExceeded"

	// ErrCodeResourceNotFound for service response error code
	// "ResourceNotFound".
	//
	// The resource being accessed was not found.
	ErrCodeResourceNotFound = "ResourceNotFound"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"ConflictException":     newErrorConflictException,
	"ResourceInUse":         newErrorResourceInUse,
	"ResourceLimitExceeded": newErrorResourceLimitExceeded,
	"ResourceNotFound":      newErrorResourceNotFound,
}

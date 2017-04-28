// Copyright 2017 Pulumi, Inc. All rights reserved.

package apigateway

import (
	"github.com/pulumi/coconut/pkg/resource/idl"
)

// An Amazon API Gateway (API Gateway) API resource.
type Resource struct {
	idl.NamedResource
	// If you want to create a child resource, the parent resource.  For resources without a parent, specify
	// the RestAPI's root resource.
	Parent *Resource `coco:"parent,replaces"`
	// A path name for the resource.
	PathPart string `coco:"pathPart,replaces"`
	// The RestAPI resource in which you want to create this resource.
	RestAPI *RestAPI `coco:"restAPI,replaces"`
}

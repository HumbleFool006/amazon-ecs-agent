// Copyright 2014-2015 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//	http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package ecs

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/internal/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go/internal/signer/v4"
)

// Amazon EC2 Container Service (Amazon ECS) is a highly scalable, fast, container
// management service that makes it easy to run, stop, and manage Docker containers
// on a cluster of Amazon EC2 instances. Amazon ECS lets you launch and stop
// container-enabled applications with simple API calls, allows you to get the
// state of your cluster from a centralized service, and gives you access to
// many familiar Amazon EC2 features like security groups, Amazon EBS volumes,
// and IAM roles.
//
// You can use Amazon ECS to schedule the placement of containers across your
// cluster based on your resource needs, isolation policies, and availability
// requirements. Amazon EC2 Container Service eliminates the need for you to
// operate your own cluster management and configuration management systems
// or worry about scaling your management infrastructure.
type ECS struct {
	*aws.Service
}

// Used for custom service initialization logic
var initService func(*aws.Service)

// Used for custom request initialization logic
var initRequest func(*aws.Request)

// New returns a new ECS client.
func New(config *aws.Config) *ECS {
	service := &aws.Service{
		Config:       aws.DefaultConfig.Merge(config),
		ServiceName:  "ecs",
		APIVersion:   "2014-11-13",
		JSONVersion:  "1.1",
		TargetPrefix: "AmazonEC2ContainerServiceV20141113",
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(jsonrpc.Build)
	service.Handlers.Unmarshal.PushBack(jsonrpc.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(jsonrpc.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(jsonrpc.UnmarshalError)

	// Run custom service initialization if present
	if initService != nil {
		initService(service)
	}

	return &ECS{service}
}

// newRequest creates a new request for a ECS operation and runs any
// custom request initialization.
func (c *ECS) newRequest(op *aws.Operation, params, data interface{}) *aws.Request {
	req := aws.NewRequest(c.Service, op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}

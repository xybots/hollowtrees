// Copyright © 2018 Banzai Cloud
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package grpcplugin

import (
	"net"

	"github.com/goph/emperror"
	"google.golang.org/grpc"

	"github.com/banzaicloud/hollowtrees/pkg/grpcplugin/proto"
)

// Serve registers the EventHandler and starts the GRPC server
func Serve(bindAddress string, handler EventHandler, opt ...grpc.ServerOption) error {
	listener, err := net.Listen("tcp", bindAddress)
	if err != nil {
		return emperror.Wrap(err, "failed to listen")
	}

	grpcServer := grpc.NewServer(opt...)
	proto.RegisterEventHandlerServer(grpcServer, NewHandler(handler))

	return grpcServer.Serve(listener)
}

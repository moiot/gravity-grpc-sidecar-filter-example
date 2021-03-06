/*
 *
 * // Copyright 2019 , Beijing Mobike Technology Co., Ltd.
 * //
 * // Licensed under the Apache License, Version 2.0 (the "License");
 * // you may not use this file except in compliance with the License.
 * // You may obtain a copy of the License at
 * //
 * //     http://www.apache.org/licenses/LICENSE-2.0
 * //
 * // Unless required by applicable law or agreed to in writing, software
 * // distributed under the License is distributed on an "AS IS" BASIS,
 * // See the License for the specific language governing permissions and
 * // limitations under the License.
 */

package main

import (
	"github.com/hashicorp/go-plugin"
	"github.com/moiot/gravity/pkg/core"
	"github.com/moiot/gravity/pkg/filters/grpc"
)

type myFilter struct{}

func (f *myFilter) Configure(data map[string]interface{}) error {
	return nil
}

func (f *myFilter) Filter(msg *core.Msg) (bool, error) {
	for k := range msg.DmlMsg.Data {
		msg.DmlMsg.Data[k] = "hello grpc"
	}

	return true, nil
}

func (f *myFilter) Close() error {
	return nil
}

func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: grpc.HandshakeConfig,
		Plugins: map[string]plugin.Plugin{
			grpc.PluginName: &grpc.FilterGRPCPlugin{Impl: &myFilter{}},
		},
		GRPCServer: plugin.DefaultGRPCServer,
	})
}


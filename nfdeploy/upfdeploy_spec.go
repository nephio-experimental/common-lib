/*
Copyright 2022-2023 The Nephio Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package nfdeploy

import (
	resource "k8s.io/apimachinery/pkg/api/resource"
)

type InterfaceConfig struct {
	Name   string   `json:"interfaceName"`
	IpAddr []string `json:"ipAddr"`
	Vlan   []string `json:"vlan"`
}

type BgpInterface struct {
	Name       string `json:"interfaceName,omitempty"`
	NeighborIp string `json:"neighborIp,omitempty"`
}

// BgpConfig specifies parameters for BGP related configuration for UPF and SMF
type BgpConfig struct {
	VirtualRouterName   string         `json:"virtualRouterName"`
	VirtualRouterNumber int            `json:"virtualRouterNumber"`
	RouteId             string         `json:"routeId"`
	AsNumber            int            `json:"asNumber"`
	PeerAsNumber        int            `json:"peerAsNumber"`
	RouteDistinguisher  string         `json:"routeDistinguisher"`
	Interfaces          []BgpInterface `json:"interfaces"`
}

// UpfCapacity specifies the operational paramerters of the UPF. The UpfDeploy
// operator should translate them into node resource requirements.
type UpfCapacity struct {
	UplinkThroughput   resource.Quantity `json:"uplinkThroughput"`
	DownlinkThroughput resource.Quantity `json:"downlinkThroughput"`
	MaximumConnections int               `json:"maximumConnections"`
}

// UpfDeploySpec specifies config parameters for UPF
type UpfDeploySpec struct {
	BgpConfigs   []BgpConfig       `json:"BgpConfig,omitempty"`
	Capacity     UpfCapacity       `json:"capacity,omitempty"`
	N3Interfaces []InterfaceConfig `json:"N3Interfaces,omitempty"`
	N4Interfaces []InterfaceConfig `json:"N4Interfaces,omitempty"`
	N6Interfaces []InterfaceConfig `json:"N6Interfaces,omitempty"`
	N9Interfaces []InterfaceConfig `json:"N9Interfaces,omitempty"`
	VendorRef    *ObjectReference  `json:"vendorRef,omitempty"`
}

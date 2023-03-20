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

package udm

// UdmDeploySpec specifies config parameters for UDM
type UdmDeploySpec struct {
	// Machine resources needed by the NF.
	CapacityProfile CapacityProfile `json:"capacityProfile" yaml:"capacityProfile,omitempty"`

	// Metadata regarding the NF.
	NfInfo NfInfo `json:"nfInfo" yaml:"nfInfo,omitempty"`
}

type CapacityProfile struct {
	// The number of CPU cores requested by the NF.
	RequestedCpu int32 `json:"requestedCpu" yaml:"requestedCpu,omitempty"`

	// The memory requested by the NF in bytes.
	RequestedMemory int32 `json:"requestedMemory" yaml:"requestedMemory,omitempty"`
}

type NfInfo struct {
	// The name of the NF vendor.
	Vendor string `json:"vendor" yaml:"vendor,omitempty"`

	// The version of the NF.
	Version string `json:"version" yaml:"version,omitempty"`
}

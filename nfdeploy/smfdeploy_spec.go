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

// SmfDeploySpec defines the desired state of SmfDeploy
type SmfDeploySpec struct {
	BgpConfigs    []BgpConfig       `json:"BgpConfig,omitempty"`
	N4Interfaces  []InterfaceConfig `json:"N4Interfaces,omitempty"`
	N7Interfaces  []InterfaceConfig `json:"N7Interfaces,omitempty"`
	N10Interfaces []InterfaceConfig `json:"N10Interfaces,omitempty"`
	N11Interfaces []InterfaceConfig `json:"N11Interfaces,omitempty"`
	VendorRef     *ObjectReference  `json:"vendorRef,omitempty"`
}

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

import v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type NFDeploy struct {
	v1.TypeMeta   `json:",inline" yaml:",inline"`
	v1.ObjectMeta `json:"metadata,omitempty" yaml:"metadata,omitempty"`
	Spec          NFDeploySpec   `json:"spec,omitempty" yaml:"spec,omitempty"`
	Status        NFDeployStatus `json:"status,omitempty" yaml:"status,omitempty"`
}

type NFDeploySpec struct {
	Plmn  Plmn   `json:"plmn,omitempty" yaml:"plmn,omitempty"`
	Sites []Site `json:"sites,omitempty" yaml:"sites,omitempty"`
}

type Plmn struct {
	MCC int `json:"mcc,omitempty" yaml:"mcc,omitempty"`
	MNC int `json:"mnc,omitempty" yaml:"mnc,omitempty"`
}

type Site struct {
	Id           string `json:"id,omitempty" yaml:"id,omitempty"`
	LocationName string `json:"locationName,omitempty" yaml:"locationName,omitempty"`
	ClusterName  string `json:"clusterName,omitempty" yaml:"clusterName,omitempty"`
	NFType       string `json:"nfType,omitempty" yaml:"nfType,omitempty"`
	NFTypeName   string `json:"nfTypeName,omitempty" yaml:"nfTypeName,omitempty"`
	NFVendor     string `json:"nfVendor,omitempty" yaml:"nfVendor,omitempty"`
	NFVersion    string `json:"nfVersion,omitempty" yaml:"nfVersion,omitempty"`
	NFNamespace  string `json:"nfNamespace,omitempty" yaml:"nfNamespace,omitempty"`

	Connectivities []Connectivity `json:"connectivities,omitempty" yaml:"connectivities,omitempty"`
}

type Connectivity struct {
	NeighborName string `json:"neighborName,omitempty" yaml:"neighborName,omitempty"`
}

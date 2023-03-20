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
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type UPFCapacityProfile struct {
	v1.TypeMeta   `json:",inline" yaml:",inline"`
	v1.ObjectMeta `json:"metadata,omitempty" yaml:"metadata,omitempty"`
	Spec          UPFCapacityProfileSpec `json:"spec,omitempty" yaml:"spec,omitempty"`
}

type UPFCapacityProfileSpec struct {
	Throughput string `json:"throughput,omitempty" yaml:"throughput,omitempty"`
}

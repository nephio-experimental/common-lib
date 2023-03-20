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

import (
	"github.com/nephio-project/common-lib/nfdeploy"
)

// UdmDeployStatus defines the observed state of a deployed UDM instance.
type UdmDeployStatus struct {
	// The generation observed by the deployment controller.
	ObservedGeneration int32 `json:"observedGeneration"`

	// Current service state of the UDM.
	Conditions []nfdeploy.NFCondition `json:"conditions,omitempty"`
}

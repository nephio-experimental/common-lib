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
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// TODO: Reconsider the name of this ConditionType.
// The NFDeploymentConditionType and NFDeployConditionType is
// reserved for NFDeploy.
type NFConditionType string

const (
	// Reconciling implies that the deployment is progressing.
	// Reconciliation for a deployment is considered when a new version
	// is adopted, when new pods scale up or old pods scale down, or when
	// required peering is in progress.
	// Condition name follows Kpt guidelines.
	Reconciling NFConditionType = "Reconciling"

	// Deployment is unable to make progress towards Reconciliation.
	// Reasons could be Pod creation failure, Peering failure etc.
	// Condition name follows Kpt guidelines.
	Stalled NFConditionType = "Stalled"

	// The Deployment is considered available when following conditions hold:
	// 1. At-least the minimal set of Pods are up and running for at-least
	//minReadySeconds.
	// 2. The Deployment is ready for required peering.
	Available NFConditionType = "Available"

	// The Deployment is making progress towards peering on the required
	// interfaces. A successful peering implies that the NF is reachable by
	// the required peers and is a able to reach them.
	Peering NFConditionType = "Peering"

	// The Deployment is available and has peered successfully on required
	// interfaces.
	// At this stage, the deployment is ready to serve requests.
	Ready NFConditionType = "Ready"
)

type NFCondition struct {
	// Type of deployment condition.
	Type NFConditionType `json:"type"`

	// Status of the condition, one of True, False, Unknown.
	Status corev1.ConditionStatus `json:"status"`

	// The last time this condition was probed for.
	LastProbeTime metav1.Time `json:"lastProbeTime,omitempty"`

	// Last time the condition transitioned from one status to another.
	LastTransitionTime metav1.Time `json:"lastTransitionTime,omitempty"`

	// The reason for the condition's last transition.
	Reason string `json:"reason,omitempty"`

	// A human readable message indicating details about the transition.
	Message string `json:"message,omitempty"`
}

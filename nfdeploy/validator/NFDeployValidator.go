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

package nfdeployvalidator

import (
	"errors"

	"github.com/nephio-project/nf-deploy-controller/api/v1alpha1"
)

type void struct{}

// connected : Represents if two nodes are connected
var connected void

// present : Represents if a node is present in list
var present void

type NFType string

const UnspecifiedNFType NFType = "unspecified"

// ValidateNFDeploy is the validation function for validating a given NFDeploy
// An NFDeploy is invalid when
// - NFDeploy name is empty
// - Two sites (NFs) with same name are present
// - NFType is not recognised (Other than AMF/SMF/UPF)
// - More than one connection is present between two sites
// - site A is present as connection of site B, but not present in list of sites
// - site A is present as connection of site B, but site B is not present as
//
//	connection of site A
func ValidateNFDeploy(nfDeploy v1alpha1.NfDeploy) error {
	var presentNodes = make(map[string]void)
	if nfDeploy.Name == "" {
		return errors.New("NFDeploy name cannot be empty")
	}
	for _, site := range nfDeploy.Spec.Sites {
		if _, isPresent := presentNodes[site.Id]; isPresent {
			return errors.New("NF with id - " + site.Id + " is already present")
		}
		presentNodes[site.Id] = present
		if NFType(site.NFType) == UnspecifiedNFType {
			return errors.New("NFType " + site.NFType + " is unrecognised")
		}
	}
	var presentConnections = make(map[string]map[string]void)

	for _, site := range nfDeploy.Spec.Sites {
		for _, connection := range site.Connectivities {
			if _, present := presentConnections[site.Id]; !present {
				presentConnections[site.Id] = make(map[string]void)
			}
			if _, isPresent := presentConnections[site.Id][connection.NeighborName]; isPresent {
				return errors.New(
					"Multiple connections found between " + site.Id +
						" and " + connection.NeighborName,
				)
			}
			presentConnections[site.Id][connection.NeighborName] = connected
		}
	}
	for _, site := range nfDeploy.Spec.Sites {
		for _, connection := range site.Connectivities {
			if _, isPresent := presentNodes[connection.NeighborName]; !isPresent {
				return errors.New("NF with id " + connection.NeighborName + " is not present")
			}
			if _, isPresent := presentConnections[connection.NeighborName][site.Id]; !isPresent {
				return errors.New(
					"Connectivity between " + connection.NeighborName +
						" and " + site.Id + " is not present",
				)
			}
		}
	}
	return nil
}

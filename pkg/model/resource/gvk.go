/*
Copyright 2021 The Kubernetes Authors.

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

package resource

import (
	"fmt"
)

// GVK stores the Group - Version - Kind triplet that uniquely identifies a resource.
// In kubebuilder, the k8s fully qualified group is stored as Group and Domain to improve UX.
type GVK struct {
	Group   string `json:"group,omitempty"`
	Domain  string `json:"domain,omitempty"`
	Version string `json:"version"`
	Kind    string `json:"kind"`
}

// QualifiedGroup returns the fully qualified group name with the available information.
func (gvk GVK) QualifiedGroup() string {
	switch "" {
	case gvk.Domain:
		return gvk.Group
	case gvk.Group:
		return gvk.Domain
	default:
		return fmt.Sprintf("%s.%s", gvk.Group, gvk.Domain)
	}
}

// IsEqualTo compares two GVK objects.
func (gvk GVK) IsEqualTo(other GVK) bool {
	return gvk.Group == other.Group &&
		gvk.Domain == other.Domain &&
		gvk.Version == other.Version &&
		gvk.Kind == other.Kind
}

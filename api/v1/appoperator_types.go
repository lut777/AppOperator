/*


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

package v1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// AppOperatorSpec defines the desired state of AppOperator
type AppOperatorSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// corev1 "k8s.io/api/core/v1"
	// Check that resources.requests.storage must be configured
	corev1.ResourceRequirements `json:"resources"`
	// Address must be a valid URL
	Address string `json:"address"`
	// User must be with format of `a_b` (alphanumeric characters, separated by only one `_`)
	User string `json:"user,omitempty"`
	// Type can only be `full` or `partial`.
	Type string `json:"type,omitempty"`
	// StorageSize must be a valid Quantity (https://godoc.org/k8s.io/apimachinery/pkg/api/resource#Quantity)
	// and must be greater than resources.requests.storage
	StorageSize string `json:"storageSize,omitempty"`
	// MaxBackups must be >=1 && <= 10.
	MaxBackups int32 `json:"maxBackups"`

}

// AppOperatorStatus defines the observed state of AppOperator
type AppOperatorStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Path string `json:"path"`
}

// +kubebuilder:object:root=true

// AppOperator is the Schema for the appoperators API
type AppOperator struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`

	Spec   AppOperatorSpec   `json:"spec"`
	Status AppOperatorStatus `json:"status"`
}

// +kubebuilder:object:root=true

// AppOperatorList contains a list of AppOperator
type AppOperatorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AppOperator `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AppOperator{}, &AppOperatorList{})
}

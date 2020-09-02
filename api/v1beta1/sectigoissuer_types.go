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

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

func init() {
	SchemeBuilder.Register(&SectigoIssuer{}, &SectigoIssuerList{})
}

// SectigoIssuerSpec defines the desired state of SectigoIssuer
type SectigoIssuerSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// URL is the base URL for the sectigo API
	URL string `json:"url"`

	Provisioner SectigoProvisioner `json:"provisioner"`
}

// SectigoIssuerStatus defines the observed state of SectigoIssuer
type SectigoIssuerStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Conditions []SectigoIssuerCondition `json:"conditions:omitempty"`
}

// +kubebuilder:object:root=true

// SectigoIssuer is the Schema for the sectigoissuers API
// +kubebuilder:subresource:status
type SectigoIssuer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SectigoIssuerSpec   `json:"spec,omitempty"`
	Status SectigoIssuerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SectigoIssuerList contains a list of SectigoIssuer
type SectigoIssuerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SectigoIssuer `json:"items"`
}

// SecretKeySelector contains the reference to a secret.
type SecretKeySelector struct {
	// The key of the secret to select from. Must be a valid secret key.
	// +optional
	Key string `json:"key,omitempty"`
}

// SectigoProvisioner contains the configuration for requesting certificates from Sectigo
type SectigoProvisioner struct {
	// The name of the Sectigo provisioner.
	Name string `json:"name"`

	// CertTypeName is the type of certificate that you wish to use
	CertTypeName string `json:"certTypeName"`

	// The customer's login URI
	LoginURI string `json:"loginURI"`

	// Reference to the customer's organization id
	OrgIDRef SecretKeySelector `json:"orgIDRef"`

	// Reference to the customer's username
	UsernameRef SecretKeySelector `json:"usernameRef"`

	// Reference to the customer's password
	PasswordRef SecretKeySelector `json:"passwordRef"`

	// Reference to the customer's secret key
	SecretKeyRef SecretKeySelector `json:"secretKeyRef"`

	ClientCertAuthEnabled bool `json:"clientCertAuthEnabled"`

	// Reference to the customer's client cert
	ClientCertRef SecretKeySelector `json:"clientCertRef"`

	// Reference to the customer's client cert private key
	ClientCertPrivKeyRef SecretKeySelector `json:"clientCertPrivKeyRef"`
}

// ConditionType represents a SectigoIssuer condition type.
// +kubebuilder:validation:Enum=Ready
type ConditionType string

const (
	// ConditionReady indicates that a SectigoIssuer is ready for use.
	ConditionReady ConditionType = "Ready"
)

// ConditionStatus represents a conditions's status.
// +kubebuilder:validation:Enum=True;False;Unknown
type ConditionStatus string

// These are valid condition statuses. "ConditionTrue" means a resource is in
// the condition; "ConditionFalse" means a resource is not in the condition;
// "ConditionUnknown" means kubernetes can't decide if a resource is in the
// condition or not. In the future, we could add other intermediate
// conditions, e.g. ConditionDegraded.
const (
	// ConditionTrue represents the fact that a given condition is true
	ConditionTrue ConditionStatus = "True"

	// ConditionTrue represents the fact that a given condition is false
	ConditionFalse ConditionStatus = "False"

	// ConditionTrue represents the fact that a given condition is unknown
	ConditionUnknown ConditionStatus = "Unknown"
)

// SectigoIssuerCondition contains condition information for the issuer.
type SectigoIssuerCondition struct {
	// Type of condition, currently ('Ready').
	Type ConditionType `json:"type"`

	// Status of the condition, one of ('True', 'False', 'Unknown').
	// +kubebuilder:validation:Enum=True;False;Unknown
	Status ConditionStatus `json:"status"`

	// LastTransitionTime is the timestamp corresponding to the last status
	// change of this condition.
	// +optional
	LastTransitionTime *metav1.Time `json:"lastTransitionTime,omitempty"`

	// Reason is a brief machine readable explanation for the condition's last
	// transition.
	// +optional
	Reason string `json:"reason,omitempty"`

	// Message is a human readable description of the details of the last
	// transition, completing reason.
	// +optional
	Message string `json:"message,omitempty"`
}

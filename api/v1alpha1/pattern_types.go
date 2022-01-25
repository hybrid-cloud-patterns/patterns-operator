/*
Copyright 2022.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.
//  https://pkg.go.dev/encoding/json#Marshal

type PatternParameter struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of Pattern. Edit pattern_types.go to remove/update
	Name  string `json:"name"`
	Value string `json:"value"`
}

// PatternSpec defines the desired state of Pattern
type PatternSpec struct {
	// SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Name     string `json:"name"`
	SiteName string `json:"siteName"`

	GitSpec    GitSpec           `json:"gitSpec"`
	GitOpsSpec GitOpsSpec        `json:"gitOpsSpec,omitempty"`
	ImageSpec  ImageRegistrySpec `json:"imageSpec"`

	Validation     bool `json:"validation,omitempty"`
	AnonymousUsage bool `json:"anonymousUsage,omitempty"`

	Parameters []PatternParameter `json:"parameters,omitempty"`
}

type GitSpec struct {
	Hostname string `json:"hostname,omitempty"`
	Account  string `json:"account"`
	Secret   string `json:"secret,omitempty"`

	OriginRepo     string `json:"originRepo,omitempty"`
	TargetRepo     string `json:"targetRepo"`
	TargetRevision string `json:"targetRevision,omitempty"`

	ValuesDirectoryURL string `json:"valuesDirectoryURL,omitempty"`
}

type ImageRegistrySpec struct {
	Hostname string `json:"hostname,omitempty"`
	Account  string `json:"account"`
	Secret   string `json:"secret,omitempty"`
}

type InstallPlanType string

const (
	InstallAutomatic InstallPlanType = "Automatic"
	InstallManual    InstallPlanType = "Manual"
)

type GitOpsSpec struct {
	OperatorChannel string `json:"operatorChannel,omitempty"`
	OperatorSource  string `json:"operatorSource,omitempty"`
	OperatorCSV     string `json:"operatorCSV,omitempty"`

	SyncPolicy          InstallPlanType `json:"syncPolicy,omitempty"`
	InstallPlanApproval InstallPlanType `json:"installPlanApproval,omitempty"`
	UseCSV              bool            `json:"useCSV,omitempty"`
}

//global:
//  imageregistry:
//   type: quay

// PatternStatus defines the observed state of Pattern
type PatternStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Pattern is the Schema for the patterns API
type Pattern struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PatternSpec   `json:"spec,omitempty"`
	Status PatternStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// PatternList contains a list of Pattern
type PatternList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Pattern `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Pattern{}, &PatternList{})
}
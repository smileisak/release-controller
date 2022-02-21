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
    corev1 "k8s.io/api/core/v1"
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ReleaseStoreSpec defines the desired state of ReleaseStore
type ReleaseStoreSpec struct {
    Provider *ReleaseStoreProvider `json:"provider"`
}

type ReleaseStoreProvider struct {
    // +optional
    Github *GithubProvider `json:"github"`
}

type ReleaseStoreConditionType string

type ReleaseStoreStatusCondition struct {
    Type   ReleaseStoreConditionType `json:"type"`
    Status corev1.ConditionStatus    `json:"status"`

    // +optional
    Reason string `json:"reason,omitempty"`

    // +optional
    Message string `json:"message,omitempty"`

    // +optional
    LastTransitionTime metav1.Time `json:"lastTransitionTime,omitempty"`
}

// ReleaseStoreStatus defines the observed state of ReleaseStore
type ReleaseStoreStatus struct {
    // +optional
    Conditions []ReleaseStoreStatusCondition `json:"conditions"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:resource:scope=Cluster,categories={release,controller}

// ReleaseStore is the Schema for the releasestores API
type ReleaseStore struct {
    metav1.TypeMeta   `json:",inline"`
    metav1.ObjectMeta `json:"metadata,omitempty"`

    Spec   ReleaseStoreSpec   `json:"spec,omitempty"`
    Status ReleaseStoreStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ReleaseStoreList contains a list of ReleaseStore
type ReleaseStoreList struct {
    metav1.TypeMeta `json:",inline"`
    metav1.ListMeta `json:"metadata,omitempty"`
    Items           []ReleaseStore `json:"items"`
}

func init() {
    SchemeBuilder.Register(&ReleaseStore{}, &ReleaseStoreList{})
}

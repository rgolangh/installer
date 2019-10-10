/*
Copyright 2018 The Kubernetes Authors.

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
	//"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// OvirtMachineProviderSpec is the type that will be embedded in a Machine.Spec.ProviderSpec field
// for an Ovirt VM. It is used by the Ovirt machine actuator to create a single machine instance.
type OvirtMachineProviderSpec struct {
	metav1.TypeMeta `json:",inline"`

	// The name of the secret containing the openstack credentials
	CloudsSecret string `json:"ovirtSecret"`

	// The flavor reference for the flavor for your server instance.
	Flavor string `json:"flavor"`
	// The name of the image to use for your server instance.
	Image string `json:"image"`

	// The ssh key to inject in the instance
	KeyName string `json:"keyName,omitempty"`

	// The machine ssh username
	SshUserName string `json:"sshUserName,omitempty"`

	// A networks object. Required parameter when there are multiple networks defined for the tenant.
	// When you do not specify the networks parameter, the server attaches to the only network created for the current tenant.
	Networks []NetworkParam `json:"networks,omitempty"`

	// The name of the secret containing the user data (startup script in most cases)
	UserDataSecret *corev1.SecretReference `json:"userDataSecret,omitempty"`

	RootVolume RootVolume `json:"root_volume,omitempty"`
}

type NetworkParam struct {
	// The UUID of the network. Required if you omit the port attribute.
	UUID string `json:"uuid,omitempty"`
	// A fixed IPv4 address for the NIC.
	FixedIp string `json:"fixed_ip,omitempty"`
	// Filters for optional network query
	Filter Filter `json:"filter,omitempty"`
}

type Filter struct {
	Status       string `json:"status,omitempty"`
	Name         string `json:"name,omitempty"`
	AdminStateUp *bool  `json:"admin_state_up,omitempty"`
	TenantID     string `json:"tenant_id,omitempty"`
	ProjectID    string `json:"project_id,omitempty"`
	Shared       *bool  `json:"shared,omitempty"`
	ID           string `json:"id,omitempty"`
	Marker       string `json:"marker,omitempty"`
	Limit        int    `json:"limit,omitempty"`
	SortKey      string `json:"sort_key,omitempty"`
	SortDir      string `json:"sort_dir,omitempty"`
	Tags         string `json:"tags,omitempty"`
	TagsAny      string `json:"tags-any,omitempty"`
	NotTags      string `json:"not-tags,omitempty"`
	NotTagsAny   string `json:"not-tags-any,omitempty"`
}

type RootVolume struct {
	VolumeType string `json:"volumeType"`
	Size       int    `json:"diskSize,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// OvirtClusterProviderSpec of an oVirt cluster
// +k8s:openapi-gen=true
type OvirtClusterProviderSpec struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// ExternalNetworkID is the ID of an external OpenStack Network. This is necessary
	// to get public internet to the VMs.
	ExternalNetworkID string `json:"externalNetworkId,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// OvirtClusterProviderStatus
// +k8s:openapi-gen=true
type OvirtClusterProviderStatus struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// CACertificate is a PEM encoded CA Certificate for the control plane nodes.
	CACertificate []byte

	// CAPrivateKey is a PEM encoded PKCS1 CA PrivateKey for the control plane nodes.
	CAPrivateKey []byte

	// Network contains all information about the created OpenStack Network.
	// It includes Subnets and Router.
	Network *Network `json:"network,omitempty"`
}

// Network
type Network struct {
	Name string `json:"name"`
	ID   string `json:"id"`

	Subnet *Subnet `json:"subnet,omitempty"`
	Router *Router `json:"router,omitempty"`
}

// Subnet
type Subnet struct {
	Name string `json:"name"`
	ID   string `json:"id"`

	CIDR string `json:"cidr"`
}

// Router represents basic information about the associated OpenStack Neutron Router
type Router struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

func init() {
	SchemeBuilder.Register(&OvirtMachineProviderSpec{})
	SchemeBuilder.Register(&OvirtClusterProviderSpec{})
	SchemeBuilder.Register(&OvirtClusterProviderStatus{})
}

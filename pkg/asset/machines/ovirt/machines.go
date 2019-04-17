// Package ovirt generates Machine objects for ovirt.
package ovirt

import (
	"fmt"

	machineapi "github.com/openshift/cluster-api/pkg/apis/machine/v1beta1"
	ovirtprovider "github.com/ovirt/cluster-api-provider-ovirt/pkg/apis/ovirtclusterproviderconfig/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"

	"github.com/openshift/installer/pkg/types"
	"github.com/openshift/installer/pkg/types/ovirt"
)

// Machines returns a list of machines for a machinepool.
func Machines(clusterID string, config *types.InstallConfig, pool *types.MachinePool, role, userDataSecret string) ([]machineapi.Machine, error) {
	if configPlatform := config.Platform.Name(); configPlatform != ovirt.Name {
		return nil, fmt.Errorf("non-ovirt configuration: %q", configPlatform)
	}
	if poolPlatform := pool.Platform.Name(); poolPlatform != ovirt.Name {
		return nil, fmt.Errorf("non-ovirt machine-pool: %q", poolPlatform)
	}
	platform := config.Platform.Ovirt

	total := int64(1)
	if pool.Replicas != nil {
		total = *pool.Replicas
	}
	provider := provider(clusterID, config.Networking.MachineCIDR.String(), platform, userDataSecret)
	var machines []machineapi.Machine
	for idx := int64(0); idx < total; idx++ {
		machine := machineapi.Machine{
			TypeMeta: metav1.TypeMeta{
				APIVersion: "machine.openshift.io/v1beta1",
				Kind:       "Machine",
			},
			ObjectMeta: metav1.ObjectMeta{
				Namespace: "openshift-machine-api",
				Name:      fmt.Sprintf("%s-%s-%d", clusterID, pool.Name, idx),
				Labels: map[string]string{
					"machine.openshift.io/cluster-api-cluster":      clusterID,
					"machine.openshift.io/cluster-api-machine-role": role,
					"machine.openshift.io/cluster-api-machine-type": role,
				},
			},
			Spec: machineapi.MachineSpec{
				ProviderSpec: machineapi.ProviderSpec{
					Value: &runtime.RawExtension{Object: provider},
				},
				// we don't need to set Versions, because we control those via cluster operators.
			},
		}
		machines = append(machines, machine)
	}

	return machines, nil
}

func provider(clusterID string, networkInterfaceAddress string, platform *ovirt.Platform, userDataSecret string) *ovirtprovider.OvirtMachineProviderSpec {
	return &ovirtprovider.OvirtMachineProviderSpec{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "ovirtproviderconfig.k8s.io/v1alpha1",
			Kind:       "OvirtMachineProviderSpec",
		},
		ClusterId:  platform.ClusterID,
		TemplateId: clusterID,
	}
}

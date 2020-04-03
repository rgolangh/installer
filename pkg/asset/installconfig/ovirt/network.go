package ovirt

import (
	"fmt"
	"sort"

	ovirtsdk4 "github.com/ovirt/go-ovirt"
	"gopkg.in/AlecAivazis/survey.v1"

	"github.com/openshift/installer/pkg/types/ovirt"
)

func askNetwork(c *ovirtsdk4.Connection, p *ovirt.Platform) (*ovirtsdk4.Network, error) {
	var networkName string
	var networkByNames = make(map[string]*ovirtsdk4.Network)
	var networkNames []string
	systemService := c.SystemService()
	networksResponse, err := systemService.ClustersService().ClusterService(p.ClusterID).NetworksService().List().Send()
	if err != nil {
		return nil, err
	}
	networks, ok := networksResponse.Networks()
	if !ok {
		return nil, fmt.Errorf("there are no available networks for cluster %s", p.ClusterID)
	}

	for _, network := range networks.Slice() {
		networkByNames[network.MustName()] = network
		networkNames = append(networkNames, network.MustName())
	}
	err = survey.AskOne(&survey.Select{
		Message: "Select the oVirt network",
		Help: "The oVirt network of the deployed VMs. 'ovirtmgmt' is the default network - it is recommended " +
			"to work with a dedicated network per OpenShift cluster",
		Options: networkNames,
	},
		&networkName,
		func(ans interface{}) error {
			choice := ans.(string)
			sort.Strings(networkNames)
			i := sort.SearchStrings(networkNames, choice)
			if i == len(networkNames) || networkNames[i] != choice {
				return fmt.Errorf("invalid network %s", choice)
			}
			network, ok := networkByNames[choice]
			if !ok {
				return fmt.Errorf("cannot find a network by name %s", choice)
			}
			p.NetworkName = network.MustName()
			return nil
		})
	return networkByNames[networkName], err
}

func askVNICProfileID(c *ovirtsdk4.Connection, p *ovirt.Platform, networkID string) error {
	var profileID string
	var profilesByNames = make(map[string]*ovirtsdk4.VnicProfile)
	var profileNames []string
	systemService := c.SystemService()
	response, err := systemService.NetworksService().NetworkService(networkID).VnicProfilesService().List().Send()
	if err != nil {
		return err
	}
	profiles, ok := response.Profiles()
	if !ok {
		return fmt.Errorf("there are no available network profiles")
	}

	for _, profile := range profiles.Slice() {
		profilesByNames[profile.MustName()] = profile
		profileNames = append(profileNames, profile.MustName())
	}

	if len(profilesByNames) == 1 {
		p.VNICProfileID = profilesByNames[profileNames[0]].MustId()
		return nil
	}

	// we have multiple vnic profile for the selected network
	err = survey.AskOne(&survey.Select{
		Message: "Select the oVirt cluster VNIC Profile",
		Help:    "The oVirt VNIC profile of the VMs",
		Options: profileNames,
	},
		&profileID,
		func(ans interface{}) error {
			choice := ans.(string)
			sort.Strings(profileNames)
			i := sort.SearchStrings(profileNames, choice)
			if i == len(profileNames) || profileNames[i] != choice {
				return fmt.Errorf("invalid VNIC profile %s", choice)
			}
			profile, ok := profilesByNames[choice]
			if !ok {
				return fmt.Errorf("cannot find a VNIC profile id by the name %s", choice)
			}
			p.VNICProfileID = profile.MustId()
			return nil
		})
	return err
}

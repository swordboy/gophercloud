package bootfromvolume

import (
	"github.com/swordboy/gophercloud"
	osBFV "github.com/swordboy/gophercloud/openstack/compute/v2/extensions/bootfromvolume"
	osServers "github.com/swordboy/gophercloud/openstack/compute/v2/servers"
)

// Create requests the creation of a server from the given block device mapping.
func Create(client *gophercloud.ServiceClient, opts osServers.CreateOptsBuilder) osServers.CreateResult {
	return osBFV.Create(client, opts)
}

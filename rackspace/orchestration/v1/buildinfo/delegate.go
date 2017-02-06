package buildinfo

import (
	"github.com/swordboy/gophercloud"
	os "github.com/swordboy/gophercloud/openstack/orchestration/v1/buildinfo"
)

// Get retreives build info data for the Heat deployment.
func Get(c *gophercloud.ServiceClient) os.GetResult {
	return os.Get(c)
}

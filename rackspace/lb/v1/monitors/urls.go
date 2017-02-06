package monitors

import (
	"strconv"

	"github.com/swordboy/gophercloud"
)

const (
	path        = "loadbalancers"
	monitorPath = "healthmonitor"
)

func rootURL(c *gophercloud.ServiceClient, lbID int) string {
	return c.ServiceURL(path, strconv.Itoa(lbID), monitorPath)
}

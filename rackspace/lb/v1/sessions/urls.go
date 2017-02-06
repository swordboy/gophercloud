package sessions

import (
	"strconv"

	"github.com/swordboy/gophercloud"
)

const (
	path   = "loadbalancers"
	spPath = "sessionpersistence"
)

func rootURL(c *gophercloud.ServiceClient, id int) string {
	return c.ServiceURL(path, strconv.Itoa(id), spPath)
}

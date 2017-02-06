package quotasets

import (
	"testing"

	th "github.com/swordboy/gophercloud/testhelper"
	"github.com/swordboy/gophercloud/testhelper/client"
)

func TestGetURL(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	c := client.ServiceClient()

	th.CheckEquals(t, c.Endpoint+"os-quota-sets/wat", getURL(c, "wat"))
}

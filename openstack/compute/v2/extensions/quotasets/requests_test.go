package quotasets

import (
	th "github.com/swordboy/gophercloud/testhelper"
	"github.com/swordboy/gophercloud/testhelper/client"
	"testing"
)

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetSuccessfully(t)
	actual, err := Get(client.ServiceClient(), FirstTenantID).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &FirstQuotaSet, actual)
}

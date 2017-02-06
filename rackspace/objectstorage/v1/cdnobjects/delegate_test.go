package cdnobjects

import (
	"testing"

	os "github.com/swordboy/gophercloud/openstack/objectstorage/v1/objects"
	th "github.com/swordboy/gophercloud/testhelper"
	fake "github.com/swordboy/gophercloud/testhelper/client"
)

func TestDeleteCDNObject(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	os.HandleDeleteObjectSuccessfully(t)

	res := Delete(fake.ServiceClient(), "testContainer", "testObject", nil)
	th.AssertNoErr(t, res.Err)

}

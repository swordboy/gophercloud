// +build acceptance

package v1

import (
	"testing"

	"github.com/swordboy/gophercloud"
	osServiceAssets "github.com/swordboy/gophercloud/openstack/cdn/v1/serviceassets"
	"github.com/swordboy/gophercloud/swordboy/cdn/v1/serviceassets"
	th "github.com/swordboy/gophercloud/testhelper"
)

func TestServiceAsset(t *testing.T) {
	client := newClient(t)

	t.Log("Creating Service")
	loc := testServiceCreate(t, client, "test-site-2")
	t.Logf("Created service at location: %s", loc)

	t.Log("Deleting Service Assets")
	testServiceAssetDelete(t, client, loc)
}

func testServiceAssetDelete(t *testing.T, client *gophercloud.ServiceClient, url string) {
	deleteOpts := osServiceAssets.DeleteOpts{
		All: true,
	}
	err := serviceassets.Delete(client, url, deleteOpts).ExtractErr()
	th.AssertNoErr(t, err)
	t.Log("Successfully deleted all Service Assets")
}

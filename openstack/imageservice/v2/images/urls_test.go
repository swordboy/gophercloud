package images

import (
	"testing"

	"github.com/swordboy/gophercloud"
	th "github.com/swordboy/gophercloud/testhelper"
)

const endpoint = "http://localhost:57909/v2/"

func endpointClient() *gophercloud.ServiceClient {
	return &gophercloud.ServiceClient{Endpoint: endpoint}
}

func TestListURL(t *testing.T) {
	th.AssertEquals(t, endpoint+"images", listURL(endpointClient()))
}

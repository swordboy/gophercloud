package common

import (
	"github.com/swordboy/gophercloud"
	"github.com/swordboy/gophercloud/testhelper/client"
)

const TokenID = client.TokenID

func ServiceClient() *gophercloud.ServiceClient {
	return client.ServiceClient()
}

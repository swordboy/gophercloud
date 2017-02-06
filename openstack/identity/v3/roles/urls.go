package roles

import "github.com/swordboy/gophercloud"

func listAssignmentsURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("role_assignments")
}

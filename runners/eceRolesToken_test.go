package main

import (
	"os"
	"testing"
)

type test struct {
	name string
	roles []string
	valid bool
}

// TestGetRolesToken tests GetRolesToken function for allocators, directors, proxy roles
// Did not write test case for invalid runner role as ECE returns a token for invalid runner roles as well.
func TestGetRolesToken(t *testing.T) {
	eceCoordinatorHost = os.Getenv("ECE_COORDINATOR_HOST")
	eceUser = os.Getenv("ECE_USER")
	ecePassword = os.Getenv("ECE_PASSWORD")

	var tests = []test{
		{
			"AllocatorToken",
			[]string{"allocator"},
			true,
		},
		{
			"ProxyToken",
			[]string{"proxy"},
			true,
		},
		{
			"DirectorToken",
			[]string{"director"},
			true,
		},
	}

	for _, c := range tests {
		t.Run(c.name, func(t *testing.T) {
			_, err := GetRolesToken(c.roles)
			if err != nil && c.valid {
				t.Error(c.name, " Failed", err)
			}
		})
	}

}

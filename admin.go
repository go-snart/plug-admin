// Package admin contains a plugin that provides admin-only commands for Snart Bots.
package admin

import "github.com/go-snart/snart"

// Admin is a Plug that provides admin-only functionality.
type Admin struct {
	Errs chan error
}

// New creates a usable Admin Plug.
func New() snart.Plug {
	return &Admin{
		Errs: nil,
	}
}

// String returns a codename for the Plug.
func (a *Admin) String() string {
	return "admin"
}

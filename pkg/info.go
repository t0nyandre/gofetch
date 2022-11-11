package pkg

import (
	"os"
	"os/user"
)

// Username returns the username of the logged in user in string format
func Username() string {
	current_user, err := user.Current()
	username := "Unknown"
	if err == nil {
		username = current_user.Username
	}
	return username
}

// Hostname returns the hostname of the system
func Hostname() string {
	hostname, err := os.Hostname()
	host := "Unknown"
	if err == nil {
		host = hostname
	}
	return host
}

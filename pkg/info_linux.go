package pkg

import "github.com/acobaugh/osrelease"

func OS() string {
	osrel, err := osrelease.ReadFile("/etc/os-release")
	if err != nil {
		panic(err)
	}
	return osrel["NAME"]
}

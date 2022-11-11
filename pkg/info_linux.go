package pkg

import (
	"github.com/t0nyandre/gofetch/internal"
)

func OS() string {
	osrelease, err := internal.ReadOSRelease()
	if err != nil {
		panic(err)
	}
	return osrelease["NAME"]
}

package cmd

import (
	"fmt"

	"github.com/t0nyandre/gofetch/pkg"
)

var goascii = `
`

func Gofetch() {
	fmt.Printf(`
        %s@%s
        os: %s
    `, pkg.Username(), pkg.Hostname(), pkg.OS())
}

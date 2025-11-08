// Package eflag
package eflag

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

// SetFlagsFromEnvironment parses env variables and matches them to their
// flag counterparts
func SetFlagsFromEnvironment() (error) {
	var err error
	flag.VisitAll(func(f *flag.Flag) {
		name := strings.ToUpper(strings.ReplaceAll(f.Name, "-", "_"))
		if value, ok := os.LookupEnv(name); ok {
			err2 := flag.Set(f.Name, value)
			if err2 != nil {
				err = fmt.Errorf("failed setting flag from environment: %w", err2)
			}
		}
	})

	return err
}

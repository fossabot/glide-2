package pkg

import (
	"fmt"
	"runtime"
)

// version must be set from the contents of VERSION file by go build's
// -X main.version= option in the Makefile.
var version = "devel"

// commitSha will be the hash that the binary was built from
// and will be populated by the Makefile
var commitSha = "unknown"

// buildDate captures the time when the build happened
var buildDate = "unknown"

var FullVersion string

func init() {
	FullVersion = fmt.Sprintf(
		"%s (commit: %s, runtime: %s, buildDate: %s)",
		version,
		commitSha,
		runtime.Version(),
		buildDate,
	)
}

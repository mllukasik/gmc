package build

import (
	"embed"
	"strings"
	"time"
)

type build struct {
	Version string
	Date    time.Time
}
//go:embed build.properties
var versionFile embed.FS

func instance() build {
	version := "unknown"
	date := time.Now()

	bytes, err := versionFile.ReadFile("build.properties")
	if err != nil {
		return build{
			Version: version,
			Date:    date,
		}
	}
	content := strings.TrimSpace(string(bytes))
	for _, line := range strings.Split(content, "\n") {
		elements := strings.Split(line, "=")
		if (elements[0] == "version") {
			version = elements[1]
		} else if (elements[0] == "build_date"){
			date, _ = time.Parse(time.RFC1123Z, elements[1])
		}
	}
	return build{
		Date:    date,
		Version: version,
	}
}

var buildInstnace = instance()

func Build() build {
	return buildInstnace
}

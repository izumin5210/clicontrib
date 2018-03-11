package cbuild

import (
	"strconv"
	"strings"
	"time"
)

const (
	packageName = "github.com/izumin5210/clicontrib/cbuild"
)

var (
	// Default is the default configuration for builds.
	Default = defaultConfig()
)

// Config contains the build configurations.
type Config struct {
	Name          string
	Version       string
	GitCommit     string
	GitTag        string
	GitNearestTag string
	GitTreeState  TreeState
	BuildTime     time.Time
}

var (
	name           string
	version        string
	gitCommit      string
	gitTag         string
	gitNearestTag  string
	gitTreeState   string
	buildTimestamp string
)

func defaultConfig() Config {
	var state TreeState
	err := state.UnmarshalText([]byte(gitTreeState))
	if err != nil {
		state = TreeStateDirty
	}
	st, err := strconv.Atoi(buildTimestamp)
	if err != nil {
		st = 0
	}
	return Config{
		Name:          name,
		Version:       version,
		GitCommit:     gitCommit,
		GitTag:        gitTag,
		GitNearestTag: gitNearestTag,
		GitTreeState:  state,
		BuildTime:     time.Unix(int64(st), 0),
	}
}

// Ldflags outputs ldflags parameter for this build context.
func (c *Config) Ldflags() string {
	ldflags := make(map[string]string, 7)

	ldflags["name"] = c.Name
	ldflags["version"] = c.Version
	ldflags["gitCommit"] = c.GitCommit
	ldflags["gitTag"] = c.GitTag
	ldflags["gitNearestTag"] = c.GitNearestTag
	ldflags["gitTreeState"] = c.GitTreeState.String()
	ldflags["buildTimestamp"] = strconv.FormatInt(c.BuildTime.Unix(), 10)

	result := make([]string, 0, len(ldflags))

	for k, v := range ldflags {
		result = append(result, "-X "+packageName+"."+k+"="+v)
	}

	return strings.Join(result, " ")
}

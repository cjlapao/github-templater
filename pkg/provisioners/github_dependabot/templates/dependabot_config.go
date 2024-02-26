package template

import (
	_ "embed"
)

//go:embed dependabot.yml.tpl
var DefaultDependabotTemplate string

var (
	DependencyTypes = []string{
		"direct",
		"indirect",
		"all",
		"production",
		"development",
	}
	UpdateTypes = []string{
		"version-update:semver-major",
		"version-update:semver-minor",
		"version-update:semver-patch",
	}
	InsecureExternalCodeExecution = []string{
		"allow",
		"deny",
	}
	VersioningStrategies = []string{
		"auto",
		"increase",
		"increase-if-necessary",
		"lockfile-only",
		"widen",
	}
	PackageEcosystems = []string{
		"bundler",
		"cargo",
		"composer",
		"devcontainers",
		"docker",
		"elm",
		"gitsubmodule",
		"github-actions",
		"gomod",
		"gradle",
		"maven",
		"mix",
		"npm",
		"nuget",
		"pip",
		"pub",
		"swift",
		"terraform",
	}
)

var ScheduleDays = []string{
	"monday",
	"tuesday",
	"wednesday",
	"thursday",
	"friday",
	"saturday",
	"sunday",
}

var ScheduleIntervals = []string{
	"daily",
	"weekly",
	"monthly",
}

var GroupUpdateTypes = []string{
	"major",
	"minor",
	"patch",
}

var PullRequestBranchNameSeparators = []string{
	"/",
	"-",
	"_",
}

var RebaseStrategies = []string{
	"auto",
	"disabled",
}

type DependabotConfig struct {
	Version                       int      `json:"version" yaml:"version"`
	Timezone                      string   `json:"timezone" yaml:"timezone"`
	DependencyType                string   `json:"dependency-type" yaml:"dependency-type"`
	UpdateTypes                   []string `json:"update-types" yaml:"update_-types"`
	InsecureExternalCodeExecution string   `json:"insecure-external-code-execution" yaml:"insecure-external-code-execution"`
	VersioningStrategy            string   `json:"versioning-strategy" yaml:"versioning-strategy"`
	PackageEcosystem              string   `json:"package-ecosystem" yaml:"package-ecosystem"`
	ScheduleDay                   string   `json:"schedule-day" yaml:"schedule-day"`
	ScheduleInterval              string   `json:"schedule-interval" yaml:"schedule-interval"`
}

type DependabotUpdate struct {
	Allow                         []DependabotAllow       `json:"allow,omitempty" yaml:"allow,omitempty"`
	Assignees                     []string                `json:"assignees,omitempty" yaml:"assignees,omitempty"`
	CommitMessage                 DependabotCommitMessage `json:"commit-message,omitempty" yaml:"commit-message,omitempty"`
	Directory                     string                  `json:"directory,omitempty" yaml:"directory,omitempty"`
	Groups                        []DependabotGroup       `json:"groups,omitempty" yaml:"groups,omitempty"`
	Ignore                        []DependabotIgnore      `json:"ignore,omitempty" yaml:"ignore,omitempty"`
	InsecureExternalCodeExecution string                  `json:"insecure-external-code-execution,omitempty" yaml:"insecure-external-code-execution,omitempty"`
	Labels                        []string                `json:"labels,omitempty" yaml:"labels,omitempty"`
	Milestone                     int                     `json:"milestone,omitempty" yaml:"milestone,omitempty"`
	OpenPullRequestsLimit         int                     `json:"open-pull-requests-limit,omitempty" yaml:"open-pull-requests-limit,omitempty"`
	PackageEcosystem              string                  `json:"package-ecosystem,omitempty" yaml:"package-ecosystem,omitempty"`
	PullRequestBranchName         string                  `json:"pull-request-branch-name,omitempty" yaml:"pull-request-branch-name,omitempty"`
	RebaseStrategy                string                  `json:"rebase-strategy,omitempty" yaml:"rebase-strategy,omitempty"`
}

func NewDependabotUpdate() DependabotUpdate {
	return DependabotUpdate{
		Allow:                 []DependabotAllow{},
		Assignees:             []string{},
		CommitMessage:         NewDependabotCommitMessage(),
		Directory:             "/",
		Groups:                []DependabotGroup{},
		Ignore:                []DependabotIgnore{},
		Labels:                []string{"dependencies"},
		OpenPullRequestsLimit: 5,
	}
}

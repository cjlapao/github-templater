package template

import (
	_ "embed"
	"fmt"

	"github.com/cjlapao/github-templater/pkg/diagnostics"
)

//go:embed stale.yml.tpl
var DefaultStaleTemplate string

type StaleConfigTemplate struct {
	DaysUntilStale  int         `json:"daysUntilStale,omitempty" yaml:"daysUntilStale,omitempty"`
	DaysUntilClose  interface{} `json:"daysUntilClose,omitempty" yaml:"daysUntilClose,omitempty"`
	OnlyLabels      []string    `json:"onlyLabels,omitempty" yaml:"onlyLabels,omitempty"`
	ExemptLabels    []string    `json:"exemptLabels,omitempty" yaml:"exemptLabels,omitempty"`
	ExemptProjects  bool        `json:"exemptProjects,omitempty" yaml:"exemptProjects,omitempty"`
	ExemptAssignees bool        `json:"exemptAssignees,omitempty" yaml:"exemptAssignees,omitempty"`
	StaleLabel      string      `json:"staleLabel,omitempty" yaml:"staleLabel,omitempty"`
	MarkComment     interface{} `json:"markComment,omitempty" yaml:"markComment,omitempty"`
	UnmarkComment   interface{} `json:"unmarkComment,omitempty" yaml:"unmarkComment,omitempty"`
	CloseComment    interface{} `json:"closeComment,omitempty" yaml:"closeComment,omitempty"`
	LimitPerRun     int         `json:"limitPerRun,omitempty" yaml:"limitPerRun,omitempty"`
	Only            string      `json:"only,omitempty" yaml:"only,omitempty"`
}

type StaleConfig struct {
	DaysUntilStale  int                  `json:"-" yaml:"-"`
	DaysUntilClose  int                  `json:"-" yaml:"-"`
	OnlyLabels      []string             `json:"-" yaml:"-"`
	ExemptLabels    []string             `json:"-" yaml:"-"`
	ExemptProjects  bool                 `json:"-" yaml:"-"`
	ExemptAssignees bool                 `json:"-" yaml:"-"`
	StaleLabel      string               `json:"-" yaml:"-"`
	MarkComment     string               `json:"-" yaml:"-"`
	UnmarkComment   string               `json:"-" yaml:"-"`
	CloseComment    string               `json:"-" yaml:"-"`
	LimitPerRun     int                  `json:"-" yaml:"-"`
	Only            string               `json:"-" yaml:"-"`
	Template        StaleConfigTemplate  `json:"template,omitempty,inline" yaml:"template,omitempty,inline"`
	Issues          *StaleConfigTemplate `json:"issues,omitempty" yaml:"issues,omitempty"`
	Pulls           *StaleConfigTemplate `json:"pulls,omitempty" yaml:"pulls,omitempty"`
}

func NewStaleConfig() StaleConfig {
	return StaleConfig{
		DaysUntilStale:  60,
		DaysUntilClose:  7,
		OnlyLabels:      []string{},
		ExemptLabels:    []string{},
		ExemptProjects:  false,
		ExemptAssignees: false,
		StaleLabel:      "wontfix",
		MarkComment:     "This issue has been automatically marked as stale because it has not had recent activity.\nIt will be closed if no further activity occurs.\nThank you for your contributions.",
		UnmarkComment:   "This issue has been automatically unmarked as stale because it has had recent activity.\nIt will be closed if no further activity occurs.\nThank you for your contributions.",
		CloseComment:    "This issue has been automatically closed due to lack of activity.\nThis is a bot generated comment, please do not respond.",
		LimitPerRun:     30,
		Only:            "",
	}
}

func (s StaleConfig) Validate() diagnostics.Diagnostics {
	diag := diagnostics.New()

	if s.DaysUntilStale < 0 {
		diag.AddError(fmt.Errorf("daysUntilStale must be greater than or equal to 0"))
	}

	if s.DaysUntilClose < 0 {
		diag.AddError(fmt.Errorf("daysUntilClose must be greater than or equal to 0"))
	}

	if s.LimitPerRun < 1 {
		diag.AddError(fmt.Errorf("limitPerRun must be greater than or equal to 1"))
	}

	if s.LimitPerRun > 30 {
		diag.AddError(fmt.Errorf("limitPerRun must be less than or equal to 30"))
	}

	if s.Only != "" && (s.Only != "issues" && s.Only != "pulls") {
		diag.AddError(fmt.Errorf("only must be set to issues or pulls, if not empty"))
	}

	return diag
}

func (s StaleConfig) Generate() StaleConfig {
	result := s

	result.Template = StaleConfigTemplate{
		DaysUntilStale:  s.DaysUntilStale,
		OnlyLabels:      s.OnlyLabels,
		ExemptLabels:    s.ExemptLabels,
		ExemptProjects:  s.ExemptProjects,
		ExemptAssignees: s.ExemptAssignees,
		StaleLabel:      s.StaleLabel,
		LimitPerRun:     s.LimitPerRun,
		Only:            s.Only,
	}
	if s.DaysUntilClose > 0 {
		result.Template.DaysUntilClose = s.DaysUntilClose
	} else {
		result.Template.DaysUntilClose = false
	}

	if s.MarkComment != "" {
		result.Template.MarkComment = s.MarkComment
	} else {
		result.Template.MarkComment = false
	}

	if s.UnmarkComment != "" {
		result.Template.UnmarkComment = s.UnmarkComment
	} else {
		result.Template.UnmarkComment = false
	}

	if s.CloseComment != "" {
		result.Template.CloseComment = s.CloseComment
	} else {
		result.Template.CloseComment = false
	}

	result.Issues = s.Issues
	result.Pulls = s.Pulls

	return result
}

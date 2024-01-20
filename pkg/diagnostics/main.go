package diagnostics

import "errors"

type Diagnostics struct {
	module   string
	errors   []DiagnosticItem
	warnings []DiagnosticItem
	infos    []DiagnosticItem
}

func New() Diagnostics {
	return Diagnostics{
		errors:   []DiagnosticItem{},
		warnings: []DiagnosticItem{},
		infos:    []DiagnosticItem{},
	}
}

func NewModuleDiagnostics(module string) Diagnostics {
	return Diagnostics{
		module:   module,
		errors:   []DiagnosticItem{},
		warnings: []DiagnosticItem{},
		infos:    []DiagnosticItem{},
	}
}

func (d *Diagnostics) AddError(err error) {
	item := NewErrorDiagnosticItem(d.module, err)
	d.errors = append(d.errors, item)
}

func (d *Diagnostics) AddWarning(message string) {
	item := NewDiagnosticItem(d.module, message, DiagnosticItemSeverityWarning)
	d.warnings = append(d.warnings, item)
}

func (d *Diagnostics) AddInfo(message string) {
	item := NewDiagnosticItem(d.module, message, DiagnosticItemSeverityInfo)
	d.infos = append(d.infos, item)
}

func (d *Diagnostics) Errors() []error {
	result := []error{}
	for _, item := range d.errors {
		if item.error != nil {
			result = append(result, item.error)
		} else {
			result = append(result, errors.New(item.Message))
		}
	}

	return result
}

func (d *Diagnostics) Warnings() []string {
	result := []string{}
	for _, item := range d.warnings {
		result = append(result, item.String())
	}
	return result
}

func (d *Diagnostics) Infos() []string {
	result := []string{}
	for _, item := range d.infos {
		result = append(result, item.String())
	}
	return result
}

func (d *Diagnostics) HasErrors() bool {
	return len(d.errors) > 0
}

func (d *Diagnostics) HasWarnings() bool {
	return len(d.warnings) > 0
}

func (d *Diagnostics) HasInfos() bool {
	return len(d.infos) > 0
}

func (d *Diagnostics) Append(other Diagnostics) {
	d.errors = append(d.errors, other.errors...)
	d.warnings = append(d.warnings, other.warnings...)
	d.infos = append(d.infos, other.infos...)
}

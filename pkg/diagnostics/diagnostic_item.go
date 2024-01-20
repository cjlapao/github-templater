package diagnostics

type DiagnosticItemSeverity string

const (
	DiagnosticItemSeverityError   DiagnosticItemSeverity = "error"
	DiagnosticItemSeverityWarning DiagnosticItemSeverity = "warning"
	DiagnosticItemSeverityInfo    DiagnosticItemSeverity = "info"
)

type DiagnosticItem struct {
	Module   string
	Message  string
	Severity DiagnosticItemSeverity
	error    error
}

func NewDiagnosticItem(module string, message string, severity DiagnosticItemSeverity) DiagnosticItem {
	return DiagnosticItem{
		Module:   module,
		Message:  message,
		Severity: severity,
	}
}

func NewErrorDiagnosticItem(module string, error error) DiagnosticItem {
	return DiagnosticItem{
		Module:   module,
		Message:  error.Error(),
		Severity: DiagnosticItemSeverityError,
	}
}

func (d *DiagnosticItem) Error() string {
	if d.error != nil {
		return d.error.Error()
	}

	return d.Message
}

func (d *DiagnosticItem) String() string {
	msg := ""
	if d.Module != "" {
		msg += "[" + d.Module + "] "
	}
	msg += d.Message
	return msg
}

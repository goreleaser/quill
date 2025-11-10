/*
Package event provides event types for all events that the library published onto the event bus. By convention, for each event
defined here there should be a corresponding event parser defined in the parsers/ child package.
*/
package event

import (
	"github.com/goreleaser/quill/internal"
)

const (
	typePrefix    = internal.ApplicationName
	cliTypePrefix = typePrefix + "-cli"

	TaskType = typePrefix + "-task"

	// CLIExitType is a partybus event indicating the main process is to exit
	CLIExitType         = cliTypePrefix + "-exit-event"
	CLIReportType       = cliTypePrefix + "-report"
	CLINotificationType = cliTypePrefix + "-notification"
	CLIInputPromptType  = cliTypePrefix + "-input-prompt"
)

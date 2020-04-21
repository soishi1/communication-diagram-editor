// Package source parses source codes of communication diagram.
package source

import (
	"regexp"
	"strings"
)

// Participant is a node in communication diagram.
type Participant struct {
	Name string
}

// Connection is a link between 2 participants.
type Connection struct {
	A, B Participant
}

// CommunicationType decides whether this communication is a request or a response.
type CommunicationType string

const (
	// Request is a function call, RPC, etc from source to dest.
	Request CommunicationType = "Request"
	// Response is a return value of Requests.
	Response CommunicationType = "Response"
)

// Communication is a message passed between 2 participants.
type Communication struct {
	Src, Dst Participant
	ComType  CommunicationType
	Comment  string
}

// Source represents communication diagram's definition.
type Source struct {
	Participants   []Participant
	Connections    []Connection
	Communications []Communication
}

// Parse parses the input string into Source struct.
// Returns error which describes the reason why parsing failed.
// The error string can be shown to users as-is.
func Parse(input string) (Source, error) {
	var source Source
	var err error

	lines := strings.Split(input, "\n")
	if source.Participants, lines, err = parseParticipants(lines); err != nil {
		return Source{}, err
	}
	return source, nil
}

var (
	participantRE = regexp.MustCompile(`^participant\s+([^\s]+[^\n]*[^\s]*)\s*$`)
)

func parseParticipants(lines []string) (participants []Participant, rest []string, err error) {
	for _, line := range lines {
		matches := participantRE.FindStringSubmatch(line)
		if len(matches) > 1 {
			participants = append(participants, Participant{Name: matches[1]})
		}
		rest = append(rest, line)
	}
	return participants, rest, nil
}

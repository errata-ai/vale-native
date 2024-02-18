package msg

import "encoding/json"

// Message is a struct that represents a message received from the browser.
type Message struct {
	// The command to run locally.
	Command string `json:"command"`
	// The content from a webpage's `textarea` or `contenteditable` element.
	Text string `json:"text"`
	// The markup language used by the webpage. For example, GitHub and Reddit
	// use Markdown.
	Format string `json:"format"`
	// The URL of the webpage.
	Url string `json:"url"`
}

// Response is a struct that represents a response to a message received from
// the browser.
type Response struct {
	// The command that was run locally.
	Command string `json:"command"`
	// The content from a webpage's `textarea` or `contenteditable` element.
	Input string `json:"input"`
	// The JSON-formatted output from Vale.
	Output string `json:"output"`
}

// ValeError is a struct that represents a single error from Vale.
type ValeError struct {
	Line int    `json:"line"`
	Path string `json:"path"`
	Text string `json:"text"`
	Code string `json:"code"`
	Span int    `json:"span"`
}

// NewEncodedValeError creates a new ValeError and returns its JSON-encoded
// representation.
//
// We use this function to create a Vale-style errors that originate from this
// host (rather than the CLI itself).
func NewEncodedValeError(text string) (string, error) {
	v := ValeError{Line: 0, Path: "", Text: text, Code: "", Span: 0}

	encoded, err := json.Marshal(v)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

// ValeErrorFromJSON creates a ValeError from its JSON-encoded representation.
func ValeErrorFromJSON(data string) (ValeError, error) {
	var v ValeError
	err := json.Unmarshal([]byte(data), &v)
	if err != nil {
		return ValeError{}, err
	}
	return v, nil
}

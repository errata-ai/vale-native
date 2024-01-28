package cli

import "strings"

func extForFormat(format string) string {
	switch strings.ToLower(format) {
	case "markdown":
		return "md"
	case "restructuredtext":
		return "rst"
	case "asciidoc":
		return "adoc"
	default:
		return "txt"
	}
}

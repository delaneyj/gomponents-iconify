package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bell(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M521.329 0v74.121c-163.599 35.891-286.205 181.326-286.954 355.52v319.63L-1.832 953.535v55.591h1203.665v-55.591L965.625 749.271v-319.63c-.753-174.194-123.364-319.629-286.97-355.52V0H521.329zm78.661 1024.585c-48.447 0-87.743 39.223-87.743 87.672c0 48.447 39.296 87.743 87.743 87.743c48.448 0 87.744-39.296 87.744-87.743c0-48.448-39.296-87.672-87.744-87.672z"/>`),
		g.Group(children),
	)
}
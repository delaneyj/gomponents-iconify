package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ProgrammingScriptCodeCodeAngleProgrammingFileBracket(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4 10.5L.5 7L4 3.5m6 7L13.5 7L10 3.5"/>`),
		g.Group(children),
	)
}
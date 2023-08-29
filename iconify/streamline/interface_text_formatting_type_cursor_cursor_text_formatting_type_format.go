package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceTextFormattingTypeCursorCursorTextFormattingTypeFormat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.5 5H13a.5.5 0 0 1 .5.5v3a.5.5 0 0 1-.5.5H7.5m-5 0H1a.5.5 0 0 1-.5-.5v-3A.5.5 0 0 1 1 5h1.5M4 2.5h2m-2 9h2m-1-9v9"/>`),
		g.Group(children),
	)
}
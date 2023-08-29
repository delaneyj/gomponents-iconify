package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceTextFormattingSuperScriptTextFormattingSuperscriptFormat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m1 5.5l7 8m0-8l-7 8M10.5 2v-.25A1.25 1.25 0 0 1 11.75.5h0A1.25 1.25 0 0 1 13 1.75A1.58 1.58 0 0 1 12.41 3L10.5 4.5H13"/>`),
		g.Group(children),
	)
}
package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceTextFormattingSubScriptTextFormattingSubscriptFormat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m1 .5l7 8m0-8l-7 8m9.5 2.5v-.25a1.25 1.25 0 0 1 1.25-1.25h0A1.25 1.25 0 0 1 13 10.75a1.58 1.58 0 0 1-.59 1.25l-1.91 1.5H13"/>`),
		g.Group(children),
	)
}
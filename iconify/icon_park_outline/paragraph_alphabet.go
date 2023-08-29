package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ParagraphAlphabet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M24 9h18M24 19h18M6 29h36M6 39h36"/><path d="m11 9l-4 8h8l-4-8Z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m6 19l1-2m9 2l-1-2m-8 0l4-8l4 8m-8 0h8"/></g>`),
		g.Group(children),
	)
}
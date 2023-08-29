package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TerminalRectangleOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 3h22v18H1V3Zm2 2v3h18V5H3Zm18 5H3v9h18v-9ZM7 11.086l3.414 3.414L7 17.914L5.586 16.5l2-2l-2-2L7 11.086ZM12 15h6v2h-6v-2Z"/>`),
		g.Group(children),
	)
}
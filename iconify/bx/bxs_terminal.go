package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BxsTerminal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M20 4H4a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2zM6.414 15.707L5 14.293L7.293 12L5 9.707l1.414-1.414L10.121 12l-3.707 3.707zM19 16h-7v-2h7v2z" fill="currentColor"/>`),
		g.Group(children),
	)
}
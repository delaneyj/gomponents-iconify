package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileReplaceOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m14 3l-2-2H4a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h7v2l4-3l-4-3v2H4V3h10m7 7v11a2 2 0 0 1-2 2H8a2 2 0 0 1-2-2v-2h2v2h11v-9h-5V7H8v6H6V7a2 2 0 0 1 2-2h8l5 5Z"/>`),
		g.Group(children),
	)
}
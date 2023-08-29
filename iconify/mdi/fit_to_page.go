package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FitToPage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 2H4c-1.11 0-2 .89-2 2v16c0 1.11.89 2 2 2h16c1.11 0 2-.89 2-2V4c0-1.11-.89-2-2-2m-8 2l3 3h-2v2h-2V7H9m-2 8l-3-3l3-3v2h2v2H7m5 7l-3-3h2v-2h2v2h2m2-2v-2h-2v-2h2V9l3 3"/>`),
		g.Group(children),
	)
}
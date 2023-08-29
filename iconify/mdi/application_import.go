package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ApplicationImport(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 12h9.8L8.3 9.5l1.4-1.4l4.9 4.9l-4.9 4.9l-1.4-1.4l2.5-2.5H1v-2M21 2H3c-1.1 0-2 .9-2 2v6.1h2V6h18v14H3v-4H1v4c0 1.1.9 2 2 2h18c1.1 0 2-.9 2-2V4c0-1.1-.9-2-2-2"/>`),
		g.Group(children),
	)
}
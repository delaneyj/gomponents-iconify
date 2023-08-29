package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignDirectionPlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 20h.09a5.5 5.5 0 0 0 .72 2H9a2 2 0 0 1 2-2v-8H3.5L6 9.5L3.5 7H11V3l1-1l1 1v4h5l2.5 2.5L18 12h-5m5 3v3h-3v2h3v3h2v-3h3v-2h-3v-3Z"/>`),
		g.Group(children),
	)
}
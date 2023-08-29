package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WatchlistRtl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M19 3H3v2h16zm0 6h-6v2h6zm0 6h-8v2h8zm-8-4.24H7.15L5.5 7l-1.65 3.76H0l3 3.17l-.9 4.05l3.4-2.14L8.9 18L8 13.95Z"/>`),
		g.Group(children),
	)
}
package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoWikiText(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M16 3v2h1v10l2 2V3zM9 5V3H5l2 2zM1 1L0 2l1 1v14h3v-2H3V5l2 2v10h4v-2H7V9l6 6h-2v2h4l3 3l1-1zm12 10l2 2V3h-4v2h2z"/>`),
		g.Group(children),
	)
}
package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QuotesRtl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M11 9v7h7V9c0-2.2-1.8-5-4-5h-2l1 2s2 0 2 3zM2 9v7h7V9c0-2.2-1.8-5-4-5H3l1 2s2 0 2 3z"/>`),
		g.Group(children),
	)
}
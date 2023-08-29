package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func History(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M9 6v5h.06l2.48 2.47l1.41-1.41L11 10.11V6z"/><path fill="currentColor" d="M10 1a9 9 0 0 0-7.85 13.35L.5 16H6v-5.5l-2.38 2.38A7 7 0 1 1 10 17v2a9 9 0 0 0 0-18z"/>`),
		g.Group(children),
	)
}
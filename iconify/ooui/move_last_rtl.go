package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoveLastRtl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 1h2v18H3zm13.5 1.5L15 1l-9 9l9 9l1.5-1.5L9 10z"/>`),
		g.Group(children),
	)
}
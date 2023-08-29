package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumbersFive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 4H7v7a2 2 0 0 0 2 2h6v5H7v2h8a2 2 0 0 0 2-2v-5a2 2 0 0 0-2-2H9V6h8V4Z"/>`),
		g.Group(children),
	)
}
package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumbersSeven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 4h10.5v3.113a2 2 0 0 1-.52 1.346l-4.48 4.928V20h-2v-6.613a2 2 0 0 1 .52-1.346l4.48-4.928V6H7V4Z"/>`),
		g.Group(children),
	)
}
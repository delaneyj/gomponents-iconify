package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextformatItalic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.5 3H18v2h-3.67l-2.625 14H16.5v2H6v-2h3.67l2.625-14H7.5V3Z"/>`),
		g.Group(children),
	)
}
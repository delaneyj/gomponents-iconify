package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AddAndSubtract(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 1v8h8v2h-8v8h-2v-8H3V9h8V1h2ZM3 20h18v2H3v-2Z"/>`),
		g.Group(children),
	)
}
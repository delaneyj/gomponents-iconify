package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderAll(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 4h17v17H3V4m1 1v7h7V5H4m15 7V5h-7v7h7M4 20h7v-7H4v7m15 0v-7h-7v7h7Z"/>`),
		g.Group(children),
	)
}
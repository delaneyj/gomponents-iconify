package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Menu(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 8V7h17v1H3m17 4v1H3v-1h17M3 17h17v1H3v-1Z"/>`),
		g.Group(children),
	)
}
package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewQuilt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 5v6h11V5m-5 13h5v-6h-5M4 18h5V5H4m6 13h5v-6h-5v6Z"/>`),
		g.Group(children),
	)
}
package flat_color_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TreeStructure(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#90CAF9" d="m36.9 13.8l-1.8-3.6L7.5 24l27.6 13.8l1.8-3.6L16.5 24z"/><path fill="#D81B60" d="M6 18h12v12H6z"/><path fill="#2196F3" d="M30 6h12v12H30zm0 24h12v12H30z"/>`),
		g.Group(children),
	)
}
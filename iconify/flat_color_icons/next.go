package flat_color_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Next(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#2196F3" d="M17.1 5L14 8.1L29.9 24L14 39.9l3.1 3.1L36 24z"/>`),
		g.Group(children),
	)
}
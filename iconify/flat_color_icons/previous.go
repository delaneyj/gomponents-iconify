package flat_color_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Previous(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#2196F3" d="m30.9 43l3.1-3.1L18.1 24L34 8.1L30.9 5L12 24z"/>`),
		g.Group(children),
	)
}
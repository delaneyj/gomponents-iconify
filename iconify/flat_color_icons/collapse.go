package flat_color_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Collapse(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#2196F3" d="M5 30.9L8.1 34L24 18.1L39.9 34l3.1-3.1L24 12z"/>`),
		g.Group(children),
	)
}
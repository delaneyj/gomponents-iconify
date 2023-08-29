package flat_color_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EmptyFilter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#FFCC80" d="M29 23H19L7 9h34zm0 15l-10 6V23h10zM41.5 9h-35C5.7 9 5 8.3 5 7.5S5.7 6 6.5 6h35c.8 0 1.5.7 1.5 1.5S42.3 9 41.5 9z"/>`),
		g.Group(children),
	)
}
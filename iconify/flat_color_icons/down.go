package flat_color_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Down(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="#3F51B5"><path d="M24 44L12.3 30h23.4z"/><path d="M20 6h8v27h-8z"/></g>`),
		g.Group(children),
	)
}
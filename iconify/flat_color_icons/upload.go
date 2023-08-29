package flat_color_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Upload(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="#009688"><path d="M24 10.9L35 24H13zM20 40h8v4h-8zm0-6h8v4h-8z"/><path d="M20 21h8v11h-8zM6 4h36v4H6z"/></g>`),
		g.Group(children),
	)
}
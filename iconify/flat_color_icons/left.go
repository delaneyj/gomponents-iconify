package flat_color_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Left(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="#3F51B5"><path d="m4 24l14-11.7v23.4z"/><path d="M15 20h27v8H15z"/></g>`),
		g.Group(children),
	)
}
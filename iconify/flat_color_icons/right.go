package flat_color_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Right(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="#3F51B5"><path d="M44 24L30 35.7V12.3z"/><path d="M6 20h27v8H6z"/></g>`),
		g.Group(children),
	)
}
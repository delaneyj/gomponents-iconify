package flat_color_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Up(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="#3F51B5"><path d="m24 4l11.7 14H12.3z"/><path d="M20 15h8v27h-8z"/></g>`),
		g.Group(children),
	)
}
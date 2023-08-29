package flat_color_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DownLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#3F51B5" d="m4 29l14-11.7v23.4z"/><path fill="#3F51B5" d="M42 21V8h-8v13c0 2.2-1.8 4-4 4H13v8h17c6.6 0 12-5.4 12-12z"/>`),
		g.Group(children),
	)
}
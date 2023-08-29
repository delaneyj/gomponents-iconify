package flat_color_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UpLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#3F51B5" d="m4 19l14 11.7V7.3z"/><path fill="#3F51B5" d="M42 27v13h-8V27c0-2.2-1.8-4-4-4H13v-8h17c6.6 0 12 5.4 12 12z"/>`),
		g.Group(children),
	)
}
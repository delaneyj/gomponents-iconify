package flat_color_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LeftUpTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#3F51B5" d="m19 4l11.7 14H7.3z"/><path fill="#3F51B5" d="M27 42h13v-8H27c-2.2 0-4-1.8-4-4V13h-8v17c0 6.6 5.4 12 12 12z"/>`),
		g.Group(children),
	)
}
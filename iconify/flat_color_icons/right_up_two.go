package flat_color_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RightUpTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#3F51B5" d="M29 4L17.3 18h23.4z"/><path fill="#3F51B5" d="M21 42H8v-8h13c2.2 0 4-1.8 4-4V13h8v17c0 6.6-5.4 12-12 12z"/>`),
		g.Group(children),
	)
}
package flat_color_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RightDownTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#3F51B5" d="M29 44L17.3 30h23.4z"/><path fill="#3F51B5" d="M21 6H8v8h13c2.2 0 4 1.8 4 4v17h8V18c0-6.6-5.4-12-12-12z"/>`),
		g.Group(children),
	)
}
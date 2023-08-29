package flat_color_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Menu(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#607D8B" d="M6 22h36v4H6zm0-12h36v4H6zm0 24h36v4H6z"/>`),
		g.Group(children),
	)
}
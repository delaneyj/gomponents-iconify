package flat_color_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func List(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#2196F3" d="M6 22h4v4H6zm0-8h4v4H6zm0 16h4v4H6zM6 6h4v4H6zm0 32h4v4H6zm8-16h28v4H14zm0-8h28v4H14zm0 16h28v4H14zm0-24h28v4H14zm0 32h28v4H14z"/>`),
		g.Group(children),
	)
}
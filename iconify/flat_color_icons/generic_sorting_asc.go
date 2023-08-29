package flat_color_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GenericSortingAsc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#2196F3" d="M6 6h4v4H6zm0 8h12v4H6zm0 8h20v4H6zm0 8h28v4H6zm0 8h36v4H6z"/>`),
		g.Group(children),
	)
}
package flat_color_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OrgUnit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#90CAF9" d="M10 10v28h28V10H10zm24 24H14V14h20v20z"/><path fill="#D81B60" d="M6 6h12v12H6z"/><path fill="#2196F3" d="M30 6h12v12H30zM6 30h12v12H6zm24 0h12v12H30z"/>`),
		g.Group(children),
	)
}
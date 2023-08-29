package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ControllerGenOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 18q2.5 0 4.25-1.75T18 12q0-2.5-1.75-4.25T12 6Q9.5 6 7.75 7.75T6 12q0 2.5 1.75 4.25T12 18Zm0-2q-1.65 0-2.825-1.175T8 12q0-1.65 1.175-2.825T12 8q1.65 0 2.825 1.175T16 12q0 1.65-1.175 2.825T12 16Zm0-2.6l2.8-2.8l-1.4-1.4l-2.8 2.8l1.4 1.4ZM3 21V3h18v18H3Zm2-2h14V5H5v14ZM5 5v14V5Z"/>`),
		g.Group(children),
	)
}
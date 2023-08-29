package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PanoramaHorizontalOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 17.65q1.95-.575 3.963-.862T12 16.5q2.025 0 4.038.288T20 17.65V6.375q-1.95.575-3.963.85T12 7.5q-2.025 0-4.038-.275T4 6.375V17.65ZM12 12Zm-9.975 8.5V3.475q1.875.875 4.5 1.45t5.5.575q2.875 0 5.5-.575t4.5-1.45V20.5q-1.875-.875-4.5-1.437t-5.5-.563q-2.875 0-5.5.563t-4.5 1.437Z"/>`),
		g.Group(children),
	)
}
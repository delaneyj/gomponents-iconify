package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PanoramaHorizontalSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2.025 20.5V3.475q1.875.875 4.5 1.45t5.5.575q2.875 0 5.5-.575t4.5-1.45V20.5q-1.875-.875-4.5-1.437t-5.5-.563q-2.875 0-5.5.563t-4.5 1.437Z"/>`),
		g.Group(children),
	)
}
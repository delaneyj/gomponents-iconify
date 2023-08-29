package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SwitchesSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 17q-2.075 0-3.538-1.463T2 12q0-2.075 1.463-3.538T7 7q1.25 0 2.263.55T11 9h11v6H11q-.725.9-1.738 1.45T7 17Zm4.9-4H20v-2h-8.1q.05.225.075.5T12 12q0 .225-.025.5t-.075.5Z"/>`),
		g.Group(children),
	)
}
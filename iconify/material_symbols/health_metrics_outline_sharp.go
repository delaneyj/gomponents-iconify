package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HealthMetricsOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 22v-5H2V7h5V2h10v5h5v10h-5v5H7ZM4 11h5.525l1.175 1.75l1.35-4.3h1.775l1.7 2.55H20V9h-5V4H9v5H4v2Zm5 9h6v-5h5v-2h-5.55l-1.15-1.75l-1.35 4.325h-1.8L8.45 13H4v2h5v5Zm3-8Z"/>`),
		g.Group(children),
	)
}
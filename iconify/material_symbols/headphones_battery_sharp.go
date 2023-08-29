package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeadphonesBatterySharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 18v-6q0-2.5 1.75-4.25T8 6q2.5 0 4.25 1.75T14 12v6h-4v-5h2.5v-1q0-1.875-1.313-3.188T8 7.5q-1.875 0-3.188 1.313T3.5 12v1H6v5H2Zm14 0V7h2V6h2v1h2v11h-6Z"/>`),
		g.Group(children),
	)
}
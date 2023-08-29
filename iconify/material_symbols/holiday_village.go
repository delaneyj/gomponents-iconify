package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HolidayVillage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 20V10l6-6l6 6v10H9v-5H7v5H2Zm5-7h2v-2H7v2Zm9 7V9.175L10.825 4h2.825L18 8.35V20h-2Zm4 0V7.525L16.475 4H19.3L22 6.7V20h-2Z"/>`),
		g.Group(children),
	)
}
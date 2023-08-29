package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HolidayVillageOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 20V10l6-6l6 6v10H2Zm2-2h3v-3h2v3h3v-7.175l-4-4l-4 4V18Zm3-5v-2h2v2H7Zm9 7V9.175L10.825 4h2.825L18 8.35V20h-2Zm4 0V7.525L16.475 4H19.3L22 6.7V20h-2ZM4 18h8h-8Z"/>`),
		g.Group(children),
	)
}
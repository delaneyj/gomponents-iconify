package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SetTopStackSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="none"><path d="M2.5 2A1.5 1.5 0 0 0 1 3.5v5A1.5 1.5 0 0 0 2.5 10h7A1.5 1.5 0 0 0 11 8.5v-5A1.5 1.5 0 0 0 9.5 2h-7zm3.92 1.729l.634.983l1.04.2a.5.5 0 0 1 .25.853l-.776.734l.41 1.356a.5.5 0 0 1-.745.568L6 7.645l-1.233.778a.5.5 0 0 1-.746-.568l.41-1.356l-.775-.734a.5.5 0 0 1 .25-.854l1.04-.199l.634-.983a.5.5 0 0 1 .84 0z" fill="currentColor"/><path d="M3.085 11A1.5 1.5 0 0 0 4.5 12H10a3 3 0 0 0 3-3V5.5a1.5 1.5 0 0 0-1-1.415V9a2 2 0 0 1-2 2H3.085z" fill="currentColor"/><path d="M5.085 13A1.5 1.5 0 0 0 6.5 14h4A4.5 4.5 0 0 0 15 9.5v-2a1.5 1.5 0 0 0-1-1.415V9.5a3.5 3.5 0 0 1-3.5 3.5H5.085z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}
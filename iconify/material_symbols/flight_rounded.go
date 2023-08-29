package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlightRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.5 21.725q-.425.125-.713-.1t-.287-.65q0-.2.113-.413t.262-.337L10.5 19v-5.5l-7.225 2.125q-.525.15-.9-.125T2 14.65q0-.275.162-.563t.413-.412L10.5 9V3.5q0-.65.425-1.075T12 2q.65 0 1.075.425T13.5 3.5V9l7.925 4.675q.25.125.413.425t.162.575q0 .55-.375.825t-.9.125L13.5 13.5V19l1.625 1.225q.15.125.263.338t.112.412q0 .425-.288.65t-.712.1L12 21l-2.5.725Z"/>`),
		g.Group(children),
	)
}
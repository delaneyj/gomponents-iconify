package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QuickreplyRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 18h-1q-.425 0-.713-.288T17 17v-4q0-.425.288-.713T18 12h2.875q.4 0 .613.338t.062.712L20.3 16h.975q.425 0 .65.35t.025.725l-2.475 4.975q-.1.175-.288.138T19 21.95V18ZM6 18l-2.3 2.3q-.475.475-1.088.213T2 19.575V4q0-.825.588-1.413T4 2h16q.825 0 1.413.588T22 4v5q0 .425-.288.713T21 10h-4q-.825 0-1.413.588T15 12v5q0 .425-.288.713T14 18H6Z"/>`),
		g.Group(children),
	)
}
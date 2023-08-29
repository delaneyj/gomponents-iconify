package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PrintError(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 21v-4H2v-6q0-1.275.875-2.138T5 8h14q1 0 1.763.563T21.825 10H18q-.825 0-1.413.588T16 12v3H8v4h8v2H6ZM6 7V3h12v4H6Zm13 14q-.425 0-.713-.288T18 20q0-.425.288-.713T19 19q.425 0 .713.288T20 20q0 .425-.288.713T19 21Zm-1-4v-5h2v5h-2Z"/>`),
		g.Group(children),
	)
}
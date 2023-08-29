package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VariablesRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 17q-.425 0-.713-.288T19 16v-2h-2q-.425 0-.713-.288T16 13q0-.425.288-.713T17 12h2v-2q0-.425.288-.713T20 9q.425 0 .713.288T21 10v2h2q.425 0 .713.288T24 13q0 .425-.288.713T23 14h-2v2q0 .425-.288.713T20 17ZM4 14q-.425 0-.713-.288T3 13V5q0-.425.288-.713T4 4h16q.425 0 .713.288T21 5v2h-2q-.825 0-1.413.588T17 9v1h-1q-.825 0-1.413.588T14 12v2H4Z"/>`),
		g.Group(children),
	)
}
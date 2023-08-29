package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignalCellularPauseRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 22q-.425 0-.713-.288T15 21v-5q0-.425.288-.713T16 15q.425 0 .713.288T17 16v5q0 .425-.288.713T16 22Zm4 0q-.425 0-.713-.288T19 21v-5q0-.425.288-.713T20 15q.425 0 .713.288T21 16v5q0 .425-.288.713T20 22ZM4.425 22q-.675 0-.938-.613T3.7 20.3L20.3 3.7q.475-.475 1.088-.213t.612.938V12q0 .425-.288.713T21 13h-6q-.825 0-1.413.588T13 15v6q0 .425-.288.713T12 22H4.425Z"/>`),
		g.Group(children),
	)
}
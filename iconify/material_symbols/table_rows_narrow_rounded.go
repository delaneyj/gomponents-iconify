package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableRowsNarrowRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 15.5q-.425 0-.713-.288T3 14.5V14q0-.425.288-.713T4 13h16q.425 0 .713.288T21 14v.5q0 .425-.288.713T20 15.5H4ZM4 11q-.425 0-.713-.288T3 10v-.5q0-.425.288-.713T4 8.5h16q.425 0 .713.288T21 9.5v.5q0 .425-.288.713T20 11H4Zm0-4.5q-.425 0-.713-.288T3 5.5V5q0-.425.288-.713T4 4h16q.425 0 .713.288T21 5v.5q0 .425-.288.713T20 6.5H4ZM4 20q-.425 0-.713-.288T3 19v-.5q0-.425.288-.713T4 17.5h16q.425 0 .713.288T21 18.5v.5q0 .425-.288.713T20 20H4Z"/>`),
		g.Group(children),
	)
}
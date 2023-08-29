package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoveGroupRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 18q-.825 0-1.413-.588T6 16v-1q0-.425.288-.713T7 14q.425 0 .713.288T8 15v1h12V6H8v1q0 .425-.288.713T7 8q-.425 0-.713-.288T6 7V4q0-.825.588-1.413T8 2h12q.825 0 1.413.588T22 4v12q0 .825-.588 1.413T20 18H8Zm-4 4q-.825 0-1.413-.588T2 20V7q0-.425.288-.713T3 6q.425 0 .713.288T4 7v13h13q.425 0 .713.288T18 21q0 .425-.288.713T17 22H4Zm9.175-10H7q-.425 0-.713-.288T6 11q0-.425.288-.713T7 10h6.175l-.9-.9Q12 8.825 12 8.412t.3-.712q.275-.275.7-.275t.7.275l2.6 2.6q.3.3.3.7t-.3.7l-2.6 2.6q-.275.275-.687.288T12.3 14.3q-.275-.275-.275-.7t.275-.7l.875-.9Z"/>`),
		g.Group(children),
	)
}
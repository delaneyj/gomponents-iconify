package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhonelinkOffRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.1 21.9L17.2 20H15q-.425 0-.712-.287T14 19v-2.2l-8-8V17h4.5q.625 0 1.063.438T12 18.5q0 .625-.438 1.063T10.5 20h-7q-.625 0-1.063-.438T2 18.5q0-.625.438-1.063T3.5 17H4V6.8L2.1 4.9q-.275-.275-.275-.7t.275-.7q.275-.275.7-.275t.7.275l17 17q.275.275.275.7t-.275.7q-.275.275-.7.275t-.7-.275ZM22 19l-8-7.85V9q0-.425.288-.713T15 8h6q.425 0 .713.288T22 9v10ZM8.85 6l-2-2H20q.425 0 .713.288T21 5q0 .425-.288.713T20 6H8.85Z"/>`),
		g.Group(children),
	)
}
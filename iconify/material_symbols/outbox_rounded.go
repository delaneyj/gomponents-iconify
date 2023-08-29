package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OutboxRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 21q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v14q0 .825-.588 1.413T19 21H5Zm7-5q.8 0 1.475-.413t1.1-1.087q.15-.225.375-.363t.5-.137H19V5H5v9h3.55q.275 0 .5.138t.375.362q.425.675 1.1 1.088T12 16Zm-1-6.15l-.875.875q-.3.3-.713.288T8.7 10.7q-.275-.3-.287-.7t.287-.7l2.6-2.6q.15-.15.325-.212T12 6.425q.2 0 .375.063t.325.212l2.6 2.6q.3.3.287.7t-.287.7q-.3.3-.713.313t-.712-.288L13 9.85V13q0 .425-.288.713T12 14q-.425 0-.713-.288T11 13V9.85Z"/>`),
		g.Group(children),
	)
}
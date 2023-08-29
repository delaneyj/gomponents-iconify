package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoreRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 19q-.5 0-.938-.225t-.712-.625l-3.525-5Q3.45 12.625 3.45 12t.375-1.15l3.525-5q.275-.4.713-.625T9 5h10q.825 0 1.413.588T21 7v10q0 .825-.588 1.413T19 19H9Zm.5-6q.425 0 .713-.288T10.5 12q0-.425-.288-.713T9.5 11q-.425 0-.713.288T8.5 12q0 .425.288.713T9.5 13Zm3.5 0q.425 0 .713-.288T14 12q0-.425-.288-.713T13 11q-.425 0-.713.288T12 12q0 .425.288.713T13 13Zm3.5 0q.425 0 .713-.288T17.5 12q0-.425-.288-.713T16.5 11q-.425 0-.713.288T15.5 12q0 .425.288.713T16.5 13Z"/>`),
		g.Group(children),
	)
}
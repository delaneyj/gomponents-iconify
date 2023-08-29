package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoreOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.5 13q.425 0 .713-.288T10.5 12q0-.425-.288-.713T9.5 11q-.425 0-.713.288T8.5 12q0 .425.288.713T9.5 13Zm3.5 0q.425 0 .713-.288T14 12q0-.425-.288-.713T13 11q-.425 0-.713.288T12 12q0 .425.288.713T13 13Zm3.5 0q.425 0 .713-.288T17.5 12q0-.425-.288-.713T16.5 11q-.425 0-.713.288T15.5 12q0 .425.288.713T16.5 13ZM9 19q-.5 0-.938-.225t-.712-.625L3 12l4.35-6.15q.275-.4.713-.625T9 5h10q.825 0 1.413.588T21 7v10q0 .825-.588 1.413T19 19H9Zm10-2V7v10ZM9 17h10V7H9l-3.55 5L9 17Z"/>`),
		g.Group(children),
	)
}
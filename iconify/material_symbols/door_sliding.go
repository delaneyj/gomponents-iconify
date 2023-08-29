package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoorSliding(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 13q.425 0 .713-.288T10 12q0-.425-.288-.713T9 11q-.425 0-.713.288T8 12q0 .425.288.713T9 13Zm6 0q.425 0 .713-.288T16 12q0-.425-.288-.713T15 11q-.425 0-.713.288T14 12q0 .425.288.713T15 13ZM3 21v-2h1V5q0-.825.588-1.413T6 3h5.5v16h1V3H18q.825 0 1.413.588T20 5v14h1v2H3Z"/>`),
		g.Group(children),
	)
}
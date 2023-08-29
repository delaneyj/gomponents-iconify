package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoorOpen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21v-2h2V5q0-.825.588-1.413T7 3h10q.825 0 1.413.588T19 5v14h2v2H3Zm12-2h2V5h-4.5V3.9q1.1.2 1.8 1.025T15 6.85V19Zm-4-6q.425 0 .713-.288T12 12q0-.425-.288-.713T11 11q-.425 0-.713.288T10 12q0 .425.288.713T11 13Z"/>`),
		g.Group(children),
	)
}
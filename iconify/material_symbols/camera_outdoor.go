package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CameraOutdoor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 21V9l8-6l8 6v2h-7q-.825 0-1.413.588T11 13v4q0 .825.588 1.413T13 19h7v2H4Zm9-3q-.425 0-.713-.288T12 17v-4q0-.425.288-.713T13 12h4q.425 0 .713.288T18 13v1l2-1.05v4.1L18 16v1q0 .425-.288.713T17 18h-4Z"/>`),
		g.Group(children),
	)
}
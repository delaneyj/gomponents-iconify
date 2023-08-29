package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AutoAwesomeMotionRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22q-.825 0-1.413-.588T10 20v-8q0-.825.588-1.413T12 10h8q.825 0 1.413.588T22 12v8q0 .825-.588 1.413T20 22h-8Zm-6-5V8q0-.825.588-1.413T8 6h9q.425 0 .713.288T18 7q0 .425-.288.713T17 8H8v9q0 .425-.288.713T7 18q-.425 0-.713-.288T6 17Zm-4-4V4q0-.825.588-1.413T4 2h9q.425 0 .713.288T14 3q0 .425-.288.713T13 4H4v9q0 .425-.288.713T3 14q-.425 0-.713-.288T2 13Z"/>`),
		g.Group(children),
	)
}
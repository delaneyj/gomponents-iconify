package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PregnancyRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 6q-.825 0-1.413-.588T9 4q0-.825.588-1.413T11 2q.825 0 1.413.588T13 4q0 .825-.588 1.413T11 6Zm.5 16q-.625 0-1.063-.438T10 20.5V17H9q-.425 0-.713-.288T8 16v-6q0-1.25.875-2.125T11 7q1.25 0 2.125.875T14 10q.9.375 1.45 1.2T16 13v3q0 .425-.288.713T15 17h-2v3.5q0 .625-.438 1.063T11.5 22Z"/>`),
		g.Group(children),
	)
}
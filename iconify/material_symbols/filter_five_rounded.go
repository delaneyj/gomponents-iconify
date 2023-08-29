package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilterFiveRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 15q.825 0 1.413-.588T17 13v-2q0-.825-.588-1.413T15 9h-2V7h3q.425 0 .713-.288T17 6q0-.425-.288-.713T16 5h-4q-.425 0-.713.288T11 6v4q0 .425.288.713T12 11h3v2h-3q-.425 0-.713.288T11 14q0 .425.288.713T12 15h3Zm-7 3q-.825 0-1.413-.588T6 16V4q0-.825.588-1.413T8 2h12q.825 0 1.413.588T22 4v12q0 .825-.588 1.413T20 18H8Zm-4 4q-.825 0-1.413-.588T2 20V7q0-.425.288-.713T3 6q.425 0 .713.288T4 7v13h13q.425 0 .713.288T18 21q0 .425-.288.713T17 22H4Z"/>`),
		g.Group(children),
	)
}
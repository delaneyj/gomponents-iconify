package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LibraryAddRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 14q.425 0 .713-.288T15 13v-2h2q.425 0 .713-.288T18 10q0-.425-.288-.713T17 9h-2V7q0-.425-.288-.713T14 6q-.425 0-.713.288T13 7v2h-2q-.425 0-.713.288T10 10q0 .425.288.713T11 11h2v2q0 .425.288.713T14 14Zm-6 4q-.825 0-1.413-.588T6 16V4q0-.825.588-1.413T8 2h12q.825 0 1.413.588T22 4v12q0 .825-.588 1.413T20 18H8Zm-4 4q-.825 0-1.413-.588T2 20V7q0-.425.288-.713T3 6q.425 0 .713.288T4 7v13h13q.425 0 .713.288T18 21q0 .425-.288.713T17 22H4Z"/>`),
		g.Group(children),
	)
}
package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EventListOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 21q-.825 0-1.413-.588T14 19v-4q0-.825.588-1.413T16 13h4q.825 0 1.413.588T22 15v4q0 .825-.588 1.413T20 21h-4Zm0-2h4v-4h-4v4ZM3 18q-.425 0-.713-.288T2 17q0-.425.288-.713T3 16h7q.425 0 .713.288T11 17q0 .425-.288.713T10 18H3Zm13-7q-.825 0-1.413-.588T14 9V5q0-.825.588-1.413T16 3h4q.825 0 1.413.588T22 5v4q0 .825-.588 1.413T20 11h-4Zm0-2h4V5h-4v4ZM3 8q-.425 0-.713-.288T2 7q0-.425.288-.713T3 6h7q.425 0 .713.288T11 7q0 .425-.288.713T10 8H3Zm15 9Zm0-10Z"/>`),
		g.Group(children),
	)
}
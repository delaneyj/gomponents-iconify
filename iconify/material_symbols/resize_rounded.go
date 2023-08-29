package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ResizeRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 13v-2h2v2H3Zm0-4V7h2v2H3Zm4-4V3h2v2H7Zm4 16v-2h2v2h-2Zm0-16V3h2v2h-2Zm4 16v-2h2v2h-2Zm4-4v-2h2v2h-2Zm0-4v-2h2v2h-2Zm1-4q-.425 0-.713-.288T19 8V5h-3q-.425 0-.713-.288T15 4q0-.425.288-.713T16 3h3q.825 0 1.413.588T21 5v3q0 .425-.288.713T20 9ZM5 21q-.825 0-1.413-.588T3 19v-3q0-.425.288-.713T4 15q.425 0 .713.288T5 16v3h3q.425 0 .713.288T9 20q0 .425-.288.713T8 21H5Zm14 0v-2h2q0 .825-.588 1.413T19 21ZM3 5q0-.825.588-1.413T5 3v2H3Z"/>`),
		g.Group(children),
	)
}
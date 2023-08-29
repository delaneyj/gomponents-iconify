package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MemoryOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 15q-.425 0-.713-.288T9 14v-4q0-.425.288-.713T10 9h4q.425 0 .713.288T15 10v4q0 .425-.288.713T14 15h-4Zm0 6q-.425 0-.713-.288T9 20v-1H7q-.825 0-1.413-.588T5 17v-2H4q-.425 0-.713-.288T3 14q0-.425.288-.713T4 13h1v-2H4q-.425 0-.713-.288T3 10q0-.425.288-.713T4 9h1V7q0-.825.588-1.413T7 5h2V4q0-.425.288-.713T10 3q.425 0 .713.288T11 4v1h2V4q0-.425.288-.713T14 3q.425 0 .713.288T15 4v1h2q.825 0 1.413.588T19 7v2h1q.425 0 .713.288T21 10q0 .425-.288.713T20 11h-1v2h1q.425 0 .713.288T21 14q0 .425-.288.713T20 15h-1v2q0 .825-.588 1.413T17 19h-2v1q0 .425-.288.713T14 21q-.425 0-.713-.288T13 20v-1h-2v1q0 .425-.288.713T10 21Zm-3-4h10V7H7v10Zm4-4h2v-2h-2v2Zm1-1Z"/>`),
		g.Group(children),
	)
}
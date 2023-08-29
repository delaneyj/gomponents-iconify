package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SixtyFpsSelectRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 14q-.825 0-1.413-.588T4 12V6q0-.825.588-1.413T6 4h4q.425 0 .713.288T11 5q0 .425-.288.713T10 6H6v2h3q.825 0 1.413.588T11 10v2q0 .825-.588 1.413T9 14H6Zm0-4v2h3v-2H6Zm9 4q-.825 0-1.413-.588T13 12V6q0-.825.588-1.413T15 4h3q.825 0 1.413.588T20 6v6q0 .825-.588 1.413T18 14h-3Zm3-2V6h-3v6h3ZM4 22q-.425 0-.713-.288T3 21v-3q0-.425.288-.713T4 17q.425 0 .713.288T5 18v3q0 .425-.288.713T4 22Zm4 0q-.425 0-.713-.288T7 21v-3q0-.425.288-.713T8 17q.425 0 .713.288T9 18v3q0 .425-.288.713T8 22Zm4 0q-.425 0-.713-.288T11 21v-3q0-.425.288-.713T12 17q.425 0 .713.288T13 18v3q0 .425-.288.713T12 22Zm4 0q-.425 0-.713-.288T15 21v-3q0-.425.288-.713T16 17h4q.425 0 .713.288T21 18v3q0 .425-.288.713T20 22h-4Z"/>`),
		g.Group(children),
	)
}
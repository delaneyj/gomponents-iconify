package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FitScreenRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 8V6h-2q-.425 0-.713-.288T17 5q0-.425.288-.713T18 4h2q.825 0 1.413.588T22 6v2q0 .425-.288.713T21 9q-.425 0-.713-.288T20 8ZM2 8V6q0-.825.588-1.413T4 4h2q.425 0 .713.288T7 5q0 .425-.288.713T6 6H4v2q0 .425-.288.713T3 9q-.425 0-.713-.288T2 8Zm18 12h-2q-.425 0-.713-.288T17 19q0-.425.288-.713T18 18h2v-2q0-.425.288-.713T21 15q.425 0 .713.288T22 16v2q0 .825-.588 1.413T20 20ZM4 20q-.825 0-1.413-.588T2 18v-2q0-.425.288-.713T3 15q.425 0 .713.288T4 16v2h2q.425 0 .713.288T7 19q0 .425-.288.713T6 20H4Zm2-6v-4q0-.825.588-1.413T8 8h8q.825 0 1.413.588T18 10v4q0 .825-.588 1.413T16 16H8q-.825 0-1.413-.588T6 14Z"/>`),
		g.Group(children),
	)
}
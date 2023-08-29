package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FourGMobiledataBadgeOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21q-.825 0-1.413-.588T1 19V5q0-.825.588-1.413T3 3h18q.825 0 1.413.588T23 5v14q0 .825-.588 1.413T21 21H3Zm0-2h18V5H3v14Zm0 0V5v14Zm12-2h3q.825 0 1.413-.588T20 15v-3q0-.425-.288-.713T19 11h-1.5q-.425 0-.713.288T16.5 12q0 .425.288.713T17.5 13h.5v2h-3V9h5q0-.825-.588-1.413T18 7h-3q-.825 0-1.413.588T13 9v6q0 .825.588 1.413T15 17Zm-7-3v2q0 .425.288.713T9 17q.425 0 .713-.288T10 16v-2h1q.425 0 .713-.288T12 13q0-.425-.288-.713T11 12h-1V8q0-.425-.288-.713T9 7q-.425 0-.713.288T8 8v4H6V8q0-.425-.288-.713T5 7q-.425 0-.713.288T4 8v5q0 .425.288.713T5 14h3Z"/>`),
		g.Group(children),
	)
}
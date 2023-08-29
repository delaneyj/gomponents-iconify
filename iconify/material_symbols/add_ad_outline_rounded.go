package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AddAdOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 17h-2q-.425 0-.713-.288T15 16q0-.425.288-.713T16 15h2v-2q0-.425.288-.713T19 12q.425 0 .713.288T20 13v2h2q.425 0 .713.288T23 16q0 .425-.288.713T22 17h-2v2q0 .425-.288.713T19 20q-.425 0-.713-.288T18 19v-2ZM3 21q-.825 0-1.413-.588T1 19V5q0-.825.588-1.413T3 3h14q.825 0 1.413.588T19 5v4q0 .425-.288.713T18 10q-.425 0-.713-.288T17 9V8H3v11h12q.425 0 .713.288T16 20q0 .425-.288.713T15 21H3ZM3 6h14V5H3v1Zm0 0V5v1Z"/>`),
		g.Group(children),
	)
}
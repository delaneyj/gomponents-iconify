package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlipRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 21q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h3q.425 0 .713.288T9 4q0 .425-.288.713T8 5H5v14h3q.425 0 .713.288T9 20q0 .425-.288.713T8 21H5Zm7 2q-.425 0-.713-.288T11 22V2q0-.425.288-.713T12 1q.425 0 .713.288T13 2v20q0 .425-.288.713T12 23Zm3-2v-2h2v2h-2Zm0-16V3h2v2h-2Zm4 16v-2h2q0 .825-.588 1.413T19 21Zm0-4v-2h2v2h-2Zm0-4v-2h2v2h-2Zm0-4V7h2v2h-2Zm0-4V3q.825 0 1.413.588T21 5h-2Z"/>`),
		g.Group(children),
	)
}
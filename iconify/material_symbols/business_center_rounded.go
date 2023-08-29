package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BusinessCenterRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 15v-2h2v2h-2Zm-1-9h4V4h-4v2ZM4 21q-.825 0-1.413-.588T2 19v-4h7v1q0 .425.288.713T10 17h4q.425 0 .713-.288T15 16v-1h7v4q0 .825-.588 1.413T20 21H4Zm-2-8V8q0-.825.588-1.413T4 6h4V4q0-.825.588-1.413T10 2h4q.825 0 1.413.588T16 4v2h4q.825 0 1.413.588T22 8v5h-7v-1q0-.425-.288-.713T14 11h-4q-.425 0-.713.288T9 12v1H2Z"/>`),
		g.Group(children),
	)
}
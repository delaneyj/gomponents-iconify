package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HMobiledataBadgeRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 21q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v14q0 .825-.588 1.413T19 21H5Zm5-8h4v3q0 .425.288.713T15 17q.425 0 .713-.288T16 16V8q0-.425-.288-.713T15 7q-.425 0-.713.288T14 8v3h-4V8q0-.425-.288-.713T9 7q-.425 0-.713.288T8 8v8q0 .425.288.713T9 17q.425 0 .713-.288T10 16v-3Z"/>`),
		g.Group(children),
	)
}
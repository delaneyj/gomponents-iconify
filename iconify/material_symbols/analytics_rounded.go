package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AnalyticsRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 12q-.425 0-.713.288T7 13v3q0 .425.288.713T8 17q.425 0 .713-.288T9 16v-3q0-.425-.288-.713T8 12Zm8-5q-.425 0-.713.288T15 8v8q0 .425.288.713T16 17q.425 0 .713-.288T17 16V8q0-.425-.288-.713T16 7Zm-4 7q-.425 0-.713.288T11 15v1q0 .425.288.713T12 17q.425 0 .713-.288T13 16v-1q0-.425-.288-.713T12 14Zm-7 7q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v14q0 .825-.588 1.413T19 21H5Zm7-9q.425 0 .713-.288T13 11q0-.425-.288-.713T12 10q-.425 0-.713.288T11 11q0 .425.288.713T12 12Z"/>`),
		g.Group(children),
	)
}
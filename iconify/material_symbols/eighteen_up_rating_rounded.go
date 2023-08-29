package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EighteenUpRatingRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.5 10.5v3.75q0 .325.213.537T9.25 15q.325 0 .537-.213T10 14.25v-4.5q0-.325-.213-.537T9.25 9h-1.5q-.325 0-.537.213T7 9.75q0 .325.213.537t.537.213h.75Zm4 4.5H15q.425 0 .713-.288T16 14v-4q0-.425-.288-.713T15 9h-2.5q-.425 0-.713.288T11.5 10v4q0 .425.288.713T12.5 15Zm.5-1v-1.5h1.5V14H13Zm0-2.5V10h1.5v1.5H13ZM5 21q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v14q0 .825-.588 1.413T19 21H5Z"/>`),
		g.Group(children),
	)
}
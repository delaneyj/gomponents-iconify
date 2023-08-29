package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OneKRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m13.5 12.75l1.55 1.975q.05.075.55.275q.425 0 .625-.388t-.075-.737L14.75 12l1.425-1.9q.275-.35.075-.725T15.6 9q-.175 0-.313.075t-.237.2L13.5 11.25v-1.5q0-.325-.213-.538T12.75 9q-.325 0-.537.213T12 9.75v4.5q0 .325.213.537t.537.213q.325 0 .537-.213t.213-.537v-1.5Zm-5-2.25v3.75q0 .325.213.537T9.25 15q.325 0 .537-.213T10 14.25v-4.5q0-.325-.213-.537T9.25 9h-1.5q-.325 0-.537.213T7 9.75q0 .325.213.537t.537.213h.75ZM5 21q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v14q0 .825-.588 1.413T19 21H5Z"/>`),
		g.Group(children),
	)
}
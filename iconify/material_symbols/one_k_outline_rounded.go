package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OneKOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m13.5 12.75l1.525 1.975q.05.075.575.275q.45 0 .65-.413t-.075-.762L14.75 12l1.425-1.85q.275-.35.075-.75T15.6 9q-.175 0-.325.075t-.25.2L13.5 11.25v-1.5q0-.325-.213-.538T12.75 9q-.325 0-.537.213T12 9.75v4.5q0 .325.213.537t.537.213q.325 0 .537-.213t.213-.537v-1.5Zm-5-2.25v3.75q0 .325.213.537T9.25 15q.325 0 .537-.213T10 14.25v-4.5q0-.325-.213-.537T9.25 9h-1.5q-.325 0-.537.213T7 9.75q0 .325.213.537t.537.213h.75ZM5 21q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v14q0 .825-.588 1.413T19 21H5Zm0-2h14V5H5v14ZM5 5v14V5Z"/>`),
		g.Group(children),
	)
}
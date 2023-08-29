package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TenKOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 10.5v3.75q0 .325.213.537T6.75 15q.325 0 .537-.213t.213-.537v-4.5q0-.325-.213-.537T6.75 9h-1q-.325 0-.537.213T5 9.75q0 .325.213.537t.537.213H6ZM9.5 15H12q.425 0 .713-.288T13 14v-4q0-.425-.288-.713T12 9H9.5q-.425 0-.713.288T8.5 10v4q0 .425.288.713T9.5 15Zm.5-1.5v-3h1.5v3H10Zm5.425-.75l1.525 1.975q.05.075.575.275q.45 0 .65-.413t-.075-.762L16.675 12l1.425-1.85q.275-.35.075-.75t-.65-.4q-.175 0-.325.075t-.25.2l-1.525 1.975v-1.5q0-.325-.213-.538T14.675 9q-.325 0-.537.213t-.213.537v4.5q0 .325.213.537t.537.213q.325 0 .537-.213t.213-.537v-1.5ZM5 21q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v14q0 .825-.588 1.413T19 21H5Zm0-2h14V5H5v14ZM5 5v14V5Z"/>`),
		g.Group(children),
	)
}
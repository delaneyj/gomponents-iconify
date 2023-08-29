package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TenKRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 10.5v3.75q0 .325.213.537T6.75 15q.325 0 .537-.213t.213-.537v-4.5q0-.325-.213-.537T6.75 9h-1q-.325 0-.537.213T5 9.75q0 .325.213.537t.537.213H6ZM9.5 15H12q.425 0 .713-.288T13 14v-4q0-.425-.288-.713T12 9H9.5q-.425 0-.713.288T8.5 10v4q0 .425.288.713T9.5 15Zm.5-1.5v-3h1.5v3H10Zm5.5-.75l1.55 1.975q.05.075.55.275q.425 0 .625-.388t-.075-.737L16.75 12l1.425-1.9q.275-.35.075-.725T17.6 9q-.175 0-.312.075t-.238.2L15.5 11.25v-1.5q0-.325-.212-.538T14.75 9q-.325 0-.537.213T14 9.75v4.5q0 .325.213.537t.537.213q.325 0 .537-.213t.213-.537v-1.5ZM5 21q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v14q0 .825-.588 1.413T19 21H5Z"/>`),
		g.Group(children),
	)
}
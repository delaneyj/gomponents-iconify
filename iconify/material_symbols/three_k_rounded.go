package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThreeKRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m14.5 12.75l1.55 1.975q.05.075.55.275q.425 0 .625-.388t-.075-.737L15.75 12l1.425-1.9q.275-.35.075-.725T16.6 9q-.175 0-.313.075t-.237.2L14.5 11.25v-1.5q0-.325-.213-.538T13.75 9q-.325 0-.537.213T13 9.75v4.5q0 .325.213.537t.537.213q.325 0 .537-.213t.213-.537v-1.5Zm-5 .75H7.25q-.325 0-.537.213t-.213.537q0 .325.213.537T7.25 15H10q.425 0 .713-.288T11 14v-4q0-.425-.288-.713T10 9H7.25q-.325 0-.537.213T6.5 9.75q0 .325.213.537t.537.213H9.5v1H8q-.2 0-.35.15T7.5 12q0 .2.15.35t.35.15h1.5v1ZM5 21q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v14q0 .825-.588 1.413T19 21H5Z"/>`),
		g.Group(children),
	)
}
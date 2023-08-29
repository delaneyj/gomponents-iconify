package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NineKOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.5 13.5H7.25q-.325 0-.537.213t-.213.537q0 .325.213.537T7.25 15H10q.425 0 .713-.288T11 14v-4q0-.425-.288-.713T10 9H7.5q-.425 0-.713.288T6.5 10v1.5q0 .425.288.713t.712.287h2v1Zm-1.5-2V10h1.5v1.5H8Zm6.5 1.25l1.525 1.975q.05.075.575.275q.45 0 .65-.413t-.075-.762L15.75 12l1.425-1.85q.275-.35.075-.75T16.6 9q-.175 0-.325.075t-.25.2L14.5 11.25v-1.5q0-.325-.213-.538T13.75 9q-.325 0-.537.213T13 9.75v4.5q0 .325.213.537t.537.213q.325 0 .537-.213t.213-.537v-1.5ZM5 21q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v14q0 .825-.588 1.413T19 21H5Zm0-2h14V5H5v14ZM5 5v14V5Z"/>`),
		g.Group(children),
	)
}
package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SixKOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m14.5 12.75l1.525 1.975q.05.075.575.275q.45 0 .65-.413t-.075-.762L15.75 12l1.425-1.85q.275-.35.075-.75T16.6 9q-.175 0-.325.075t-.25.2L14.5 11.25v-1.5q0-.325-.213-.538T13.75 9q-.325 0-.537.213T13 9.75v4.5q0 .325.213.537t.537.213q.325 0 .537-.213t.213-.537v-1.5ZM7.5 15H10q.425 0 .713-.288T11 14v-1.5q0-.425-.288-.713T10 11.5H8v-1h2.25q.325 0 .537-.213T11 9.75q0-.325-.213-.537T10.25 9H7.5q-.425 0-.713.288T6.5 10v4q0 .425.288.713T7.5 15Zm.5-1v-1.5h1.5V14H8Zm-3 7q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v14q0 .825-.588 1.413T19 21H5Zm0-2h14V5H5v14ZM5 5v14V5Z"/>`),
		g.Group(children),
	)
}
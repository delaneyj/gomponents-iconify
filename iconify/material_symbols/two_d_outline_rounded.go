package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwoDOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 13.5v-1h2q.425 0 .713-.288T11 11.5V10q0-.425-.288-.713T10 9H7.25q-.325 0-.537.213T6.5 9.75q0 .325.213.537t.537.213H9.5v1h-2q-.425 0-.713.288T6.5 12.5v1.75q0 .325.213.537T7.25 15h3q.325 0 .537-.213T11 14.25q0-.325-.213-.537t-.537-.213H8Zm5.75 1.5h2.75q.625 0 1.063-.438T18 13.5v-3q0-.625-.438-1.063T16.5 9h-2.75q-.325 0-.537.213T13 9.75v4.5q0 .325.213.537t.537.213Zm.75-1.5v-3h2v3h-2ZM5 21q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v14q0 .825-.588 1.413T19 21H5Zm0-2h14V5H5v14ZM5 5v14V5Z"/>`),
		g.Group(children),
	)
}
package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MpOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.5 10.5h1v2.25q0 .325.213.537t.537.213q.325 0 .537-.213T10 12.75V10.5h1v3.75q0 .325.213.537t.537.213q.325 0 .537-.213t.213-.537V10q0-.425-.288-.713T11.5 9H7q-.425 0-.713.288T6 10v4.25q0 .325.213.537T6.75 15q.325 0 .537-.213t.213-.537V10.5Zm7.5 3h2q.425 0 .713-.288T18 12.5V10q0-.425-.288-.713T17 9h-2.5q-.425 0-.713.288T13.5 10v4.25q0 .325.213.537t.537.213q.325 0 .537-.213T15 14.25v-.75Zm0-1.5v-1.5h1.5V12H15ZM5 21q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v14q0 .825-.588 1.413T19 21H5Zm0-2h14V5H5v14ZM5 5v14V5Z"/>`),
		g.Group(children),
	)
}
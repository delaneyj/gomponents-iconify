package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LuggageOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 21q-.825 0-1.413-.588T5 19V8q0-.825.588-1.413T7 6h2V3.5q0-.625.438-1.063T10.5 2h3q.625 0 1.063.438T15 3.5V6h2q.825 0 1.413.588T19 8v11q0 .825-.588 1.413T17 21q0 .425-.288.713T16 22q-.425 0-.713-.288T15 21H9q0 .425-.288.713T8 22q-.425 0-.713-.288T7 21Zm0-2h10V8H7v11Zm1.75-1q.325 0 .537-.213t.213-.537v-7.5q0-.325-.213-.537T8.75 9q-.325 0-.537.213T8 9.75v7.5q0 .325.213.537T8.75 18ZM12 18q.325 0 .537-.213t.213-.537v-7.5q0-.325-.213-.537T12 9q-.325 0-.537.213t-.213.537v7.5q0 .325.213.537T12 18Zm3.25 0q.325 0 .537-.213T16 17.25v-7.5q0-.325-.213-.537T15.25 9q-.325 0-.537.213t-.213.537v7.5q0 .325.213.537t.537.213ZM12 13.5ZM10.5 6h3V3.5h-3V6Z"/>`),
		g.Group(children),
	)
}
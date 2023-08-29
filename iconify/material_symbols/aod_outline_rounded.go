package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AodOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.75 14.5q-.325 0-.537-.213T9 13.75q0-.325.213-.537T9.75 13h4.5q.325 0 .537.213t.213.537q0 .325-.213.537t-.537.213h-4.5Zm-1-3q-.325 0-.537-.213T8 10.75q0-.325.213-.537T8.75 10h6.5q.325 0 .537.213t.213.537q0 .325-.213.537t-.537.213h-6.5ZM7 23q-.825 0-1.413-.588T5 21V3q0-.825.588-1.413T7 1h10q.825 0 1.413.588T19 3v18q0 .825-.588 1.413T17 23H7Zm0-3v1h10v-1H7Zm0-2h10V6H7v12ZM7 4h10V3H7v1Zm0 0V3v1Zm0 16v1v-1Z"/>`),
		g.Group(children),
	)
}
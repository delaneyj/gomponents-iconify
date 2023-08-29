package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExposureRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 21q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v14q0 .825-.588 1.413T19 21H5Zm0-2h14V5L5 19Zm9.5-3h-1.25q-.325 0-.537-.213t-.213-.537q0-.325.213-.537t.537-.213h1.25v-1.25q0-.325.213-.537t.537-.213q.325 0 .537.213t.213.537v1.25h1.25q.325 0 .537.213t.213.537q0 .325-.213.537T17.25 16H16v1.25q0 .325-.213.537T15.25 18q-.325 0-.537-.213t-.213-.537V16Zm-4.25-7.5q.325 0 .537-.213T11 7.75q0-.325-.213-.537T10.25 7h-3.5q-.325 0-.537.213T6 7.75q0 .325.213.537t.537.213h3.5Z"/>`),
		g.Group(children),
	)
}
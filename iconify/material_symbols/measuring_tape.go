package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MeasuringTape(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 20v-8.5q0-3.55 2.475-6.025T13.5 3q3.55 0 6.025 2.475T22 11.5q0 3.55-2.475 6.025T13.5 20H5Zm8.5-5q1.45 0 2.475-1.025T17 11.5q0-1.45-1.025-2.475T13.5 8q-1.45 0-2.475 1.025T10 11.5q0 1.45 1.025 2.475T13.5 15Zm0-2q-.625 0-1.063-.438T12 11.5q0-.625.438-1.063T13.5 10q.625 0 1.063.438T15 11.5q0 .625-.438 1.063T13.5 13ZM2 20v-5h2v5H2Z"/>`),
		g.Group(children),
	)
}
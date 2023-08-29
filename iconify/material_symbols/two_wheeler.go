package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwoWheeler(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 19q-1.65 0-2.825-1.175T0 15q0-1.425.913-2.525T3.25 11.1l-.7-.6H0V9h4.5L7 10.5L11 9h3.15L12.6 7H10V5h3.55l2.1 2.7L19 6v3h-2.3l1.75 2.3q.375-.15.763-.225T20 11q1.65 0 2.825 1.175T24 15q0 1.65-1.175 2.825T20 19q-1.65 0-2.825-1.175T16 15q0-.675.238-1.312T16.9 12.5l-.5-.6L13 17h-3l-2-1.75q-.125 1.575-1.275 2.663T4 19Zm0-2q.825 0 1.413-.588T6 15q0-.825-.588-1.413T4 13q-.825 0-1.413.588T2 15q0 .825.588 1.413T4 17Zm16 0q.825 0 1.413-.588T22 15q0-.825-.588-1.413T20 13q-.825 0-1.413.588T18 15q0 .825.588 1.413T20 17Z"/>`),
		g.Group(children),
	)
}
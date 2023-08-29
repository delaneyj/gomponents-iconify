package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CastConnectedOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 16h-3.625q-.175-.525-.388-1.038T13.55 14H16v-4h-5.675Q9.6 9.375 8.763 8.875T7.025 8H18v8Zm-6-4ZM2 20v-3q1.25 0 2.125.875T5 20H2Zm5 0q0-2.075-1.463-3.538T2 15v-2q2.925 0 4.963 2.038T9 20H7Zm4 0q0-1.875-.713-3.513t-1.925-2.85q-1.212-1.212-2.85-1.924T2 11V9q2.275 0 4.275.863t3.5 2.362q1.5 1.5 2.363 3.5T13 20h-2Zm9 0h-5q0-.5-.038-1t-.112-1H20V6H4v1.15q-.5-.075-1-.113T2 7V6q0-.825.588-1.413T4 4h16q.825 0 1.413.588T22 6v12q0 .825-.588 1.413T20 20Z"/>`),
		g.Group(children),
	)
}
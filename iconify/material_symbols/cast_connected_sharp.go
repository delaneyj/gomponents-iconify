package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CastConnectedSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 20h-7q0-.5-.038-1t-.112-1H20V6H4v1.15q-.5-.075-1-.113T2 7V4h20v16ZM2 20v-3q1.25 0 2.125.875T5 20H2Zm5 0q0-2.075-1.463-3.538T2 15v-2q2.925 0 4.963 2.038T9 20H7Zm4 0q0-1.875-.713-3.513t-1.925-2.85q-1.212-1.212-2.85-1.924T2 11V9q2.275 0 4.275.863t3.5 2.362q1.5 1.5 2.363 3.5T13 20h-2Zm3.375-4q-.875-2.725-2.788-4.8T7.026 8H18v8h-3.625Z"/>`),
		g.Group(children),
	)
}
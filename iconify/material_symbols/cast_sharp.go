package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CastSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 20v-3q1.25 0 2.125.875T5 20H2Zm5 0q0-2.075-1.463-3.538T2 15v-2q2.925 0 4.963 2.038T9 20H7Zm4 0q0-1.875-.713-3.513t-1.925-2.85q-1.212-1.212-2.85-1.924T2 11V9q2.275 0 4.275.863t3.5 2.362q1.5 1.5 2.363 3.5T13 20h-2Zm4 0q0-2.725-1.012-5.088t-2.775-4.125Q9.45 9.025 7.075 8.014T2 7V4h20v16h-7Z"/>`),
		g.Group(children),
	)
}
package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShapesSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 16q-2.925 0-4.963-2.038T2 9q0-2.925 2.038-4.963T9 2q2.925 0 4.963 2.038T16 9q0 2.925-2.038 4.963T9 16Zm-1 6v-4.05q.25.025.5.038T9 18q3.75 0 6.375-2.625T18 9q0-.25-.013-.5T17.95 8H22v14H8Z"/>`),
		g.Group(children),
	)
}
package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shapes(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 16q-2.925 0-4.963-2.038T2 9q0-2.925 2.038-4.963T9 2q2.925 0 4.963 2.038T16 9q0 2.925-2.038 4.963T9 16Zm-1 4v-2.05q.25.025.5.038T9 18q3.75 0 6.375-2.625T18 9q0-.25-.013-.5T17.95 8H20q.825 0 1.413.588T22 10v10q0 .825-.588 1.413T20 22H10q-.825 0-1.413-.588T8 20Z"/>`),
		g.Group(children),
	)
}
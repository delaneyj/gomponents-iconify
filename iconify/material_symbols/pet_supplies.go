package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PetSupplies(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.5 22q-1.475 0-2.488-1.012T6 18.5q0-.225.063-.35t-.013-.2q-.075-.075-.2-.013T5.5 18q-1.475 0-2.488-1.012T2 14.5q0-1.475 1.012-2.488T5.5 11q.575 0 1.05.15t.9.45l4.15-4.15q-.3-.425-.45-.9T11 5.5q0-1.475 1.013-2.488T14.5 2q1.475 0 2.488 1.012T18 5.5q0 .225-.063.35t.013.2q.075.075.2.013T18.5 6q1.475 0 2.488 1.012T22 9.5q0 1.475-1.012 2.488T18.5 13q-.575 0-1.05-.15t-.9-.45l-4.15 4.15q.3.425.45.9T13 18.5q0 1.475-1.012 2.488T9.5 22Z"/>`),
		g.Group(children),
	)
}
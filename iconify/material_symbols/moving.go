package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Moving(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3.4 18L2 16.6l5.3-5.3q.875-.875 2.125-.875t2.125.875l1.15 1.15q.3.3.713.3t.712-.3L18.575 8H16V6h6v6h-2V9.425l-4.475 4.45q-.875.875-2.125.875t-2.125-.875L10.1 12.7q-.275-.275-.7-.275t-.7.275L3.4 18Z"/>`),
		g.Group(children),
	)
}
package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AllMatchSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m17.275 20.25l3.475-3.45l-1.05-1.05l-2.425 2.375l-.975-.975l-1.05 1.075l2.025 2.025ZM18 23q-2.075 0-3.537-1.463T13 18q0-2.075 1.463-3.538T18 13q2.075 0 3.538 1.463T23 18q0 2.075-1.463 3.538T18 23ZM3 19l4.5-7L3 5h13.05l4.625 6.525q-.6-.25-1.388-.4T17.85 11q-2.95.05-4.9 2.138T11 18.024q0 .25.013.488t.062.487H3Z"/>`),
		g.Group(children),
	)
}
package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DryOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14.975 8q.125-1.025-.012-1.613t-.738-1.337q-.775-1-1.037-1.925T13.1 1H15q-.2 1.275.125 1.875t.9 1.4q.65.875.85 1.738T16.9 8h-1.925ZM19 8q.125-1.025-.025-1.613t-.75-1.337q-.775-1-1.025-1.925T17.125 1H19q-.2 1.275.125 1.875t.9 1.4q.65.875.85 1.738T20.9 8H19Zm0 14H2V10.975l9.625-6.25L13.25 6.35L11.3 9.5H20v2H7.7l1.925-3.1L4 12.05V20h15v2Zm-7-7v-2h10v2H12Zm0 3.5v-2h9v2h-9Zm-4-3.05Z"/>`),
		g.Group(children),
	)
}
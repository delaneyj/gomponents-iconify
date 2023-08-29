package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PrescriptionsSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m18.375 22.075l1.15-1.15l-4.45-4.45l-1.15 1.15Q13 18.55 13 19.85t.925 2.225Q14.85 23 16.15 23t2.225-.925Zm2.55-2.55l1.15-1.15Q23 17.45 23 16.15t-.925-2.225Q21.15 13 19.85 13t-2.225.925l-1.15 1.15l4.45 4.45ZM7 9h10V7H7v2Zm5-4.75q.325 0 .537-.213t.213-.537q0-.325-.213-.537T12 2.75q-.325 0-.537.213t-.213.537q0 .325.213.537T12 4.25ZM11.125 21H3V3h6.2q.325-.9 1.088-1.45T12 1q.95 0 1.713.55T14.8 3H21v8.125q-1.05-.25-2.087-.038T17 11.876V11H7v2h8.725l-2 2H7v2h4.875q-.575.875-.788 1.913T11.126 21Z"/>`),
		g.Group(children),
	)
}
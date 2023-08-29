package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PrescriptionsOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m16.95 20.65l1.15-1.15l-1.6-1.6l-1.15 1.15q-.35.35-.35.8t.35.8q.35.35.8.35t.8-.35Zm2.55-2.55l1.15-1.15q.35-.35.35-.8t-.35-.8q-.35-.35-.8-.35t-.8.35L17.9 16.5l1.6 1.6Zm-1.125 3.975Q17.45 23 16.15 23t-2.225-.925Q13 21.15 13 19.85t.925-2.225l3.7-3.7Q18.55 13 19.85 13t2.225.925Q23 14.85 23 16.15t-.925 2.225l-3.7 3.7ZM5 19V5v14Zm6.125 2H3V3h6.2q.325-.9 1.088-1.45T12 1q.95 0 1.713.55T14.8 3H21v8.125Q20.5 11 20 11t-1 .075V5H5v14h6.075Q11 19.5 11 20t.125 1ZM12 4.25q.325 0 .537-.213t.213-.537q0-.325-.213-.537T12 2.75q-.325 0-.537.213t-.213.537q0 .325.213.537T12 4.25ZM7 9V7h10v2H7Zm0 4v-2h10v.85q-.2.125-.388.288t-.387.362l-.5.5H7Zm0 4v-2h6.725L12.5 16.225q-.2.2-.362.388T11.85 17H7Z"/>`),
		g.Group(children),
	)
}
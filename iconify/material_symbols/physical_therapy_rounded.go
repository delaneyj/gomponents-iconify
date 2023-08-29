package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhysicalTherapyRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14.625 21.925q.25.1.488.063t.437-.163q.2-.125.325-.337T16 21v-3q0-.425-.288-.713T15 17H9.375q-.575 0-.938.35t-.437.8q-.075.45.138.875t.762.625l5.725 2.275ZM19.225 11q.725 0 1.25.488T21 12.7q0 .575-.325 1.05t-.875.675L18 15H9.375q-1.425 0-2.4.975T6 18.375q0 .45.125.863T6.45 20H4q-.425 0-.713-.288T3 19v-3q0-2.075 1.463-3.538T8 11h11.225ZM10 2q1.65 0 2.825 1.175T14 6q0 1.65-1.175 2.825T10 10q-1.65 0-2.825-1.175T6 6q0-1.65 1.175-2.825T10 2Z"/>`),
		g.Group(children),
	)
}
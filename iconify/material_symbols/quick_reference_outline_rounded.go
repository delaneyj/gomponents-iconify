package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QuickReferenceOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 20q.2 0 .35-.15t.15-.35v-3q0-.2-.15-.35T17 16q-.2 0-.35.15t-.15.35v3q0 .2.15.35T17 20Zm0-5q.2 0 .35-.15t.15-.35q0-.2-.15-.35T17 14q-.2 0-.35.15t-.15.35q0 .2.15.35T17 15ZM5 4v16V4v5v-5Zm3 10h2.675q.275-.575.638-1.075t.812-.925H8q-.425 0-.713.288T7 13q0 .425.288.713T8 14Zm0 4h2.075Q10 17.5 10 17t.075-1H8q-.425 0-.713.288T7 17q0 .425.288.713T8 18Zm-3 4q-.825 0-1.413-.588T3 20V4q0-.825.588-1.413T5 2h7.175q.4 0 .763.15t.637.425l4.85 4.85q.275.275.425.638t.15.762V10.3q-.475-.15-.975-.225T17 10V9h-4q-.425 0-.713-.288T12 8V4H5v16h5.675q.275.575.638 1.075t.812.925H5Zm12-10q2.075 0 3.538 1.463T22 17q0 2.075-1.463 3.538T17 22q-2.075 0-3.538-1.463T12 17q0-2.075 1.463-3.538T17 12Z"/>`),
		g.Group(children),
	)
}
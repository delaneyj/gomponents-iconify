package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QuickReferenceOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.5 20h1v-4h-1v4Zm.5-5q.2 0 .35-.15t.15-.35q0-.2-.15-.35T17 14q-.2 0-.35.15t-.15.35q0 .2.15.35T17 15ZM5 4v16V4v5v-5Zm2 10h3.675q.275-.575.638-1.075t.812-.925H7v2Zm0 4h3.075Q10 17.5 10 17t.075-1H7v2Zm-4 4V2h10l6 6v2.3q-.475-.15-.975-.225T17 10V9h-5V4H5v16h5.675q.275.575.638 1.075t.812.925H3Zm14-10q2.075 0 3.538 1.463T22 17q0 2.075-1.463 3.538T17 22q-2.075 0-3.538-1.463T12 17q0-2.075 1.463-3.538T17 12Z"/>`),
		g.Group(children),
	)
}
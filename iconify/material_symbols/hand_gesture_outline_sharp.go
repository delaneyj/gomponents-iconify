package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HandGestureOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 5.025q0-1.275-.875-2.15T18.975 2V.5q1.875 0 3.2 1.325t1.325 3.2H22ZM6 23q-2.075 0-3.538-1.463T1 18h1.5q0 1.45 1.025 2.475T6 21.5V23Zm2.475 0L1.2 12.375l1.725-1.65L7 13.575V3h2v14.425l-3.7-2.6L9.525 21H19V4h2v19H8.475ZM11 12V1h2v11h-2Zm4 0V2h2v10h-2Zm-2 4.5Z"/>`),
		g.Group(children),
	)
}
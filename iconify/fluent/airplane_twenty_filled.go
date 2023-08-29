package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AirplaneTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M9 2c.607 0 1.157.36 1.4.916l2.368 5.41l3.236-.201a1.879 1.879 0 0 1 1.992 1.758L18 10a1.879 1.879 0 0 1-1.879 1.879l-.117-.004l-3.236-.202l-2.367 5.41A1.529 1.529 0 0 1 9 18a.883.883 0 0 1-.883-.883l.002-.055l.703-5.636l-2.458-.154l-1.14 2.28A.809.809 0 0 1 4.5 14a.5.5 0 0 1-.5-.5v-2.376l-1.062-.065A1 1 0 0 1 2 10.06v-.122a1 1 0 0 1 .938-.998L4 8.875V6.5a.5.5 0 0 1 .5-.5a.81.81 0 0 1 .724.447l1.139 2.28l2.459-.154l-.698-5.58A.883.883 0 0 1 9 2Z"/>`),
		g.Group(children),
	)
}
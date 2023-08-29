package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlayLessonSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 22V2h16v9.1q-.45-.05-1-.05t-1 .05q-2.725.4-4.362 2.4T11 18q0 1.1.338 2.087T12.25 22H3Zm15 1q-2.075 0-3.538-1.463T13 18q0-2.075 1.463-3.538T18 13q2.075 0 3.538 1.463T23 18q0 2.075-1.463 3.538T18 23Zm-1.25-2.5l4-2.5l-4-2.5v5ZM7 11l2.5-1.5L12 11V4H7v7Z"/>`),
		g.Group(children),
	)
}
package uiw

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleCloseO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 0c5.523 0 10 4.477 10 10s-4.477 10-10 10S0 15.523 0 10S4.477 0 10 0Zm0 1.395a8.605 8.605 0 1 0 0 17.21a8.605 8.605 0 0 0 0-17.21Zm2.207 5.442a.682.682 0 0 1 .963.964l-2.195 2.193l2.195 2.193a.682.682 0 0 1-.963.965l-2.197-2.195l-2.195 2.195a.682.682 0 0 1-.88.071l-.084-.072a.682.682 0 0 1 0-.964l2.195-2.193l-2.195-2.193a.682.682 0 1 1 .964-.964L10.01 9.03Z"/>`),
		g.Group(children),
	)
}
package uiw

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Information(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 0c5.523 0 10 4.477 10 10s-4.477 10-10 10S0 15.523 0 10S4.477 0 10 0Zm-.145 7.21a.698.698 0 0 0-.698.697v7.558a.698.698 0 0 0 1.395 0V7.907a.698.698 0 0 0-.697-.698Zm.028-2.791a.93.93 0 1 0 0 1.86a.93.93 0 0 0 0-1.86Z"/>`),
		g.Group(children),
	)
}
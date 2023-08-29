package uiw

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InformationO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 0c5.523 0 10 4.477 10 10s-4.477 10-10 10S0 15.523 0 10S4.477 0 10 0Zm0 1.395a8.605 8.605 0 1 0 0 17.21a8.605 8.605 0 0 0 0-17.21ZM9.855 7.21c.385 0 .697.313.697.698v7.558a.698.698 0 0 1-1.395 0V7.907c0-.385.312-.698.698-.698Zm.028-2.79a.93.93 0 1 1 0 1.86a.93.93 0 0 1 0-1.86Z"/>`),
		g.Group(children),
	)
}
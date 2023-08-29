package uiw

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RightCircleO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 0c5.523 0 10 4.477 10 10s-4.477 10-10 10S0 15.523 0 10S4.477 0 10 0Zm0 1.395a8.605 8.605 0 1 0 0 17.21a8.605 8.605 0 0 0 0-17.21Zm-1.251 4.25l3.87 3.972a.682.682 0 0 1-.005.957l-3.87 3.895a.682.682 0 1 1-.967-.961l3.397-3.42l-3.402-3.49a.682.682 0 1 1 .977-.952Z"/>`),
		g.Group(children),
	)
}
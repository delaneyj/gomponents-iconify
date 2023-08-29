package uiw

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LeftCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 0c5.523 0 10 4.477 10 10s-4.477 10-10 10S0 15.523 0 10S4.477 0 10 0Zm.788 5.646l-3.87 3.971a.682.682 0 0 0 .005.957l3.87 3.895a.682.682 0 1 0 .967-.961l-3.397-3.42l3.402-3.49a.682.682 0 1 0-.977-.952Z"/>`),
		g.Group(children),
	)
}
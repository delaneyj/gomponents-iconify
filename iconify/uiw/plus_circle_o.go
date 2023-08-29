package uiw

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlusCircleO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 0c5.523 0 10 4.477 10 10s-4.477 10-10 10S0 15.523 0 10S4.477 0 10 0Zm0 1.395a8.605 8.605 0 1 0 0 17.21a8.605 8.605 0 0 0 0-17.21Zm0 4.08c.377 0 .682.305.682.681L10.68 9.28h3.124a.682.682 0 1 1 0 1.364H10.68v3.123a.682.682 0 1 1-1.363 0v-3.123H6.195a.682.682 0 1 1 0-1.363l3.123-.001V6.156c0-.376.305-.681.682-.681Z"/>`),
		g.Group(children),
	)
}
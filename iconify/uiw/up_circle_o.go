package uiw

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UpCircleO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 0c5.523 0 10 4.477 10 10s-4.477 10-10 10S0 15.523 0 10S4.477 0 10 0Zm0 1.395a8.605 8.605 0 1 0 0 17.21a8.605 8.605 0 0 0 0-17.21Zm-.48 6.241a.682.682 0 0 1 .956-.005l3.971 3.87a.682.682 0 1 1-.951.977l-3.491-3.402l-3.42 3.398a.682.682 0 1 1-.96-.968Z"/>`),
		g.Group(children),
	)
}
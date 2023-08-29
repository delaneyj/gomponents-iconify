package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CheckSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M13 .9L12 2H0v14h14V5.5l1.7-2L13 .9zM6.5 11.7L2.3 7.5l1.4-1.4l2.7 2.7L13 2.2l1.4 1.4l-7.9 8.1z"/>`),
		g.Group(children),
	)
}
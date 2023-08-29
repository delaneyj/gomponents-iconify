package uiw

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleCheck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 0c5.523 0 10 4.477 10 10s-4.477 10-10 10S0 15.523 0 10S4.477 0 10 0Zm4.922 6.583a.682.682 0 0 0-.963.053l-4.412 4.932l-2.832-3.03a.682.682 0 1 0-.996.931l3.34 3.575a.682.682 0 0 0 1.007-.01l4.91-5.489a.682.682 0 0 0-.054-.962Z"/>`),
		g.Group(children),
	)
}
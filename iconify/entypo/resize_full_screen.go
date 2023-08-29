package entypo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ResizeFullScreen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m6.987 10.987l-2.931 3.031L2 11.589V18h6.387l-2.43-2.081l3.03-2.932l-2-2zM11.613 2l2.43 2.081l-3.03 2.932l2 2l2.931-3.031L18 8.411V2h-6.387z"/>`),
		g.Group(children),
	)
}
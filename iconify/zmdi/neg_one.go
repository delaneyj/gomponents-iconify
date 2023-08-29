package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NegOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M0 171h171v42H0v-42zm320 149h-43V93l-64 22V79l101-36h6v277z"/>`),
		g.Group(children),
	)
}
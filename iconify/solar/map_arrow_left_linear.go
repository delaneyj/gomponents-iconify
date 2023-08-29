package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapArrowLeftLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="1.5" d="m19.503 20.835l-16.51-7.363c-1.324-.59-1.324-2.354 0-2.944l16.51-7.363c1.495-.667 3.047.814 2.306 2.202l-3.152 5.904c-.245.459-.245 1 0 1.458l3.152 5.904c.74 1.388-.81 2.87-2.306 2.202Z"/>`),
		g.Group(children),
	)
}
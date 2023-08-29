package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CaretLeftTwelveFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6.299 3.282C6.93 2.659 8 3.107 8 3.994v4.012c0 .887-1.07 1.335-1.701.713L4.26 6.713a1 1 0 0 1 0-1.425L6.3 3.282Z"/>`),
		g.Group(children),
	)
}
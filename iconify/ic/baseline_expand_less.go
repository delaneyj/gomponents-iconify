package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BaselineExpandLess(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 8l-6 6l1.41 1.41L12 10.83l4.59 4.58L18 14z"/>`),
		g.Group(children),
	)
}
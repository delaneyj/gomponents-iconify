package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m17.657 16.243l1.414-1.414l-7.07-7.072l-7.072 7.072l1.414 1.414L12 10.586l5.657 5.657Z"/>`),
		g.Group(children),
	)
}
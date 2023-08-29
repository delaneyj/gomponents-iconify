package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronRightDoubleS(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m17.912 11.498l-4.95 4.95l-1.414-1.414l3.535-3.536l-3.535-3.535l1.414-1.415l4.95 4.95Zm-5.461 0l-4.95 4.95l-1.414-1.414l3.535-3.536l-3.535-3.535L7.5 6.548l4.95 4.95Z"/>`),
		g.Group(children),
	)
}
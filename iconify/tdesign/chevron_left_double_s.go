package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronLeftDoubleS(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m6.088 11.498l4.95-4.95l1.414 1.415l-3.536 3.535l3.536 3.536l-1.414 1.414l-4.95-4.95Zm5.46 0l4.95-4.95l1.415 1.415l-3.536 3.535l3.536 3.536l-1.414 1.414l-4.95-4.95Z"/>`),
		g.Group(children),
	)
}
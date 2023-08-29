package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronUpDoubleS(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 6.084l4.95 4.95l-1.414 1.414l-3.535-3.536l-3.536 3.536l-1.414-1.414L12 6.084Zm0 5.46l4.95 4.95l-1.414 1.415l-3.535-3.536l-3.536 3.536l-1.414-1.415l4.95-4.95Z"/>`),
		g.Group(children),
	)
}
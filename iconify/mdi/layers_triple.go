package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LayersTriple(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 0L3 7l1.63 1.27L12 14l7.36-5.73L21 7l-9-7m7.37 10.73L12 16.47l-7.38-5.73L3 12l9 7l9-7l-1.63-1.27m0 5L12 21.47l-7.38-5.73L3 17l9 7l9-7l-1.63-1.27Z"/>`),
		g.Group(children),
	)
}
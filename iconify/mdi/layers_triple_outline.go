package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LayersTripleOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 16.54l7.37-5.74L21 12.07l-9 7l-9-7l1.62-1.26L12 16.54M12 14L3 7l9-7l9 7l-9 7m0-11.47L6.26 7L12 11.47L17.74 7L12 2.53m0 18.94l7.37-5.74L21 17l-9 7l-9-7l1.62-1.26L12 21.47"/>`),
		g.Group(children),
	)
}
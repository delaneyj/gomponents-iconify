package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LayersOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 18.54l7.37-5.74L21 14.07l-9 7l-9-7l1.62-1.26L12 18.54M12 16L3 9l9-7l9 7l-9 7m0-11.47L6.26 9L12 13.47L17.74 9L12 4.53Z"/>`),
		g.Group(children),
	)
}
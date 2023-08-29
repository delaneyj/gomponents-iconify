package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronDownS(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 15.684l-4.95-4.95L8.464 9.32L12 12.856l3.535-3.536l1.414 1.414l-4.95 4.95Z"/>`),
		g.Group(children),
	)
}
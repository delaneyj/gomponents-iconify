package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ConeTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 5.002v.5l-8.13 14.99a1 1 0 0 1-1.74 0L3 5.503v-.5C3 3.344 7.03 2 12 2s9 1.344 9 3.002"/>`),
		g.Group(children),
	)
}
package humbleicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tags(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M3 10.95V6a1 1 0 0 1 1-1h4.95a1 1 0 0 1 .707.293l7.636 7.636a1 1 0 0 1 0 1.415l-4.95 4.949a1 1 0 0 1-1.414 0l-7.636-7.636A1 1 0 0 1 3 10.948z"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="m15.636 20l5.657-5.656a1 1 0 0 0 0-1.415L13.363 5"/><circle cx="6.5" cy="8.5" r="1.5" fill="currentColor"/></g>`),
		g.Group(children),
	)
}
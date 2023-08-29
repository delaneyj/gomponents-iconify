package humbleicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M4 10.586V5a1 1 0 0 1 1-1h5.586a1 1 0 0 1 .707.293l8.5 8.5a1 1 0 0 1 0 1.414l-5.586 5.586a1 1 0 0 1-1.414 0l-8.5-8.5A1 1 0 0 1 4 10.586z"/><circle cx="8.5" cy="8.5" r="1.5" fill="currentColor"/></g>`),
		g.Group(children),
	)
}
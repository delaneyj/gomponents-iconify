package humbleicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Archive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="2"><path stroke-linejoin="round" d="M3 5a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1v3H3V5z"/><path stroke-linecap="round" d="M9.5 13h5"/><path stroke-linejoin="round" d="M4 8h16v11a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1V8z"/></g>`),
		g.Group(children),
	)
}
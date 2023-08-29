package humbleicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Calendar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="2"><path stroke-linejoin="round" d="M4 6a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1v4H4V6z"/><path stroke-linecap="round" d="M8 6.5v-3m8 3v-3"/><path stroke-linejoin="round" d="M4 10h16v9a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1v-9z"/></g>`),
		g.Group(children),
	)
}
package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SkipForward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g class="icon-tabler" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M4 6v12a2 2 0 0 0 2.75 1.84l8.25-6.1a2 2 0 0 0 0-3.5l-8.25-6.1A2 2 0 0 0 4 5.89"/><path d="M20 4v16"/></g>`),
		g.Group(children),
	)
}
package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GlobeEarth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><circle cx="12" cy="12" r="9" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/><path fill="currentColor" d="M8 6V4c2.5-1.167 8.4-2 12 4h-1a2 2 0 0 0-2 2a2 2 0 1 1-4 0a2 2 0 0 0-2-2h-1a2 2 0 0 1-2-2zm9 10h3c-1.2 1.6-3.833 3.333-5 4v-2a2 2 0 0 1 2-2zm-6 2v2c-6.4-.4-8-5.5-8-8h2a2 2 0 0 1 2 2a2 2 0 0 0 2 2a2 2 0 0 1 2 2z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 6V4c2.5-1.167 8.4-2 12 4h-1a2 2 0 0 0-2 2a2 2 0 1 1-4 0a2 2 0 0 0-2-2h-1a2 2 0 0 1-2-2zm9 10h3c-1.2 1.6-3.833 3.333-5 4v-2a2 2 0 0 1 2-2zm-6 2v2c-6.4-.4-8-5.5-8-8h2a2 2 0 0 1 2 2a2 2 0 0 0 2 2a2 2 0 0 1 2 2z"/></g>`),
		g.Group(children),
	)
}
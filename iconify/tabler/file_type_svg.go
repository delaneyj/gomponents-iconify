package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileTypeSvg(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14 3v4a1 1 0 0 0 1 1h4"/><path d="M5 12V5a2 2 0 0 1 2-2h7l5 5v4M4 20.25c0 .414.336.75.75.75H6a1 1 0 0 0 1-1v-1a1 1 0 0 0-1-1H5a1 1 0 0 1-1-1v-1a1 1 0 0 1 1-1h1.25a.75.75 0 0 1 .75.75m3-.75l2 6l2-6m6 0h-1a2 2 0 0 0-2 2v2a2 2 0 0 0 2 2h1v-3"/></g>`),
		g.Group(children),
	)
}
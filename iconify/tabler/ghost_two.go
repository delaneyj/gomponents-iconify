package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GhostTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M10 9h.01M14 9h.01M12 3a7 7 0 0 1 7 7v1h1a2 2 0 1 1 0 4h-1v3l2 3H11a6 6 0 0 1-6-5.775v-.226H4a2 2 0 0 1 0-4h1v-1a7 7 0 0 1 7-7z"/><path d="M11 14h2a1 1 0 0 0-2 0z"/></g>`),
		g.Group(children),
	)
}
package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileVector(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14 3v4a1 1 0 0 0 1 1h4M8 16.5a1.5 1.5 0 1 0 3 0a1.5 1.5 0 1 0-3 0m5-4a1.5 1.5 0 1 0 3 0a1.5 1.5 0 1 0-3 0"/><path d="M17 21H7a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h7l5 5v11a2 2 0 0 1-2 2z"/><path d="M9.5 15a2.5 2.5 0 0 1 2.5-2.5h1"/></g>`),
		g.Group(children),
	)
}
package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileTypeJpg(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14 3v4a1 1 0 0 0 1 1h4"/><path d="M5 12V5a2 2 0 0 1 2-2h7l5 5v4m-8 6h1.5a1.5 1.5 0 0 0 0-3H11v6m9-6h-1a2 2 0 0 0-2 2v2a2 2 0 0 0 2 2h1v-3M5 15h3v4.5a1.5 1.5 0 0 1-3 0"/></g>`),
		g.Group(children),
	)
}
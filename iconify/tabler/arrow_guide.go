package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowGuide(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 19a2 2 0 1 0 4 0a2 2 0 1 0-4 0m4 0h3a2 2 0 0 0 2-2V9a2 2 0 0 1 2-2h7"/><path d="m18 4l3 3l-3 3"/></g>`),
		g.Group(children),
	)
}
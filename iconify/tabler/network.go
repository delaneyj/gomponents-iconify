package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Network(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M6 9a6 6 0 1 0 12 0A6 6 0 1 0 6 9"/><path d="M12 3c1.333.333 2 2.333 2 6s-.667 5.667-2 6m0-12c-1.333.333-2 2.333-2 6s.667 5.667 2 6M6 9h12M3 19h7m4 0h7m-11 0a2 2 0 1 0 4 0a2 2 0 1 0-4 0m2-4v2"/></g>`),
		g.Group(children),
	)
}
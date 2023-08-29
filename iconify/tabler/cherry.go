package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cherry(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 16.5a3.5 3.5 0 1 0 7 0a3.5 3.5 0 1 0-7 0M14 18a3 3 0 1 0 6 0a3 3 0 1 0-6 0"/><path d="M9 13c.366-2 1.866-3.873 4.5-5.6M17 15c-1.333-2.333-2.333-5.333-1-9"/><path d="M5 6c3.667-2.667 7.333-2.667 11 0c-3.667 2.667-7.333 2.667-11 0"/></g>`),
		g.Group(children),
	)
}
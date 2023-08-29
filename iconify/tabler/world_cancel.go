package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WorldCancel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M21 12a9 9 0 1 0-8.985 9M3.6 9h16.8M3.6 15h9.9"/><path d="M11.5 3a17 17 0 0 0 0 18m1-18a16.991 16.991 0 0 1 2.53 10.275M16 19a3 3 0 1 0 6 0a3 3 0 1 0-6 0m1 2l4-4"/></g>`),
		g.Group(children),
	)
}
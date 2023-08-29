package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CarTurbine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M7 13a4 4 0 1 0 8 0a4 4 0 1 0-8 0"/><path d="M18.86 11c.088.66.14 1.512.14 2a8 8 0 1 1-8-8h6"/><path d="M11 9c2.489.108 4.489.108 6 0m0-5a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1v6a1 1 0 0 1-1 1h-2a1 1 0 0 1-1-1zm-6 9l-3.5-1.5M11 13l2.5 3m-5 0l2.5-3m0 0l3.5-1.5M11 9v4"/></g>`),
		g.Group(children),
	)
}
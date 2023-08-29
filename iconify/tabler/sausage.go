package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sausage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M5.5 5.5A2.5 2.5 0 0 0 3 8c0 7.18 5.82 13 13 13a2.5 2.5 0 1 0 0-5a8 8 0 0 1-8-8a2.5 2.5 0 0 0-2.5-2.5z"/><path d="M5.195 5.519L3.952 3.53A1 1 0 0 1 4.8 2h1.392a1 1 0 0 1 .848 1.53L5.795 5.52m12.687 12.705l1.989-1.243a1 1 0 0 1 1.53.848v1.392a1 1 0 0 1-1.53.848l-1.991-1.245"/></g>`),
		g.Group(children),
	)
}
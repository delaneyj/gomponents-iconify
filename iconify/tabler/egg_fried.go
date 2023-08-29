package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EggFried(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M9 12a3 3 0 1 0 6 0a3 3 0 1 0-6 0"/><path d="M14 3a5 5 0 0 1 4.872 6.13a3 3 0 0 1 .178 5.681a3 3 0 1 1-4.684 3.626a5 5 0 1 1-8.662-5a5 5 0 1 1 4.645-8.856A4.982 4.982 0 0 1 14 2.996z"/></g>`),
		g.Group(children),
	)
}
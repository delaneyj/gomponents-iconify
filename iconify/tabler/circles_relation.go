package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CirclesRelation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M9.183 6.117a6 6 0 1 0 4.511 3.986"/><path d="M14.813 17.883a6 6 0 1 0-4.496-3.954"/></g>`),
		g.Group(children),
	)
}
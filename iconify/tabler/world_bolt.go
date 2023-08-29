package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WorldBolt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M20.985 12.52a9 9 0 1 0-7.52 8.36M3.6 9h16.8M3.6 15h10.9"/><path d="M11.5 3a17 17 0 0 0 0 18m1-18c2.313 3.706 3.07 7.856 2.27 12M19 16l-2 3h4l-2 3"/></g>`),
		g.Group(children),
	)
}
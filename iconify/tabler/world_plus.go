package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WorldPlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M20.985 12.518a9 9 0 1 0-8.45 8.466M3.6 9h16.8M3.6 15H15"/><path d="M11.5 3a17 17 0 0 0 0 18m1-18a16.998 16.998 0 0 1 2.283 12.157M16 19h6m-3-3v6"/></g>`),
		g.Group(children),
	)
}
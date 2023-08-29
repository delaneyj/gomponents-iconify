package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mickey(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M5.5 3a3.5 3.5 0 0 1 3.25 4.8a7.017 7.017 0 0 0-2.424 2.1A3.5 3.5 0 1 1 5.5 3zm13 0a3.5 3.5 0 1 1-.826 6.902a7.013 7.013 0 0 0-2.424-2.103A3.5 3.5 0 0 1 18.5 3z"/><path d="M5 14a7 7 0 1 0 14 0a7 7 0 1 0-14 0"/></g>`),
		g.Group(children),
	)
}
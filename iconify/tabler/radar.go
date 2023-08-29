package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Radar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M21 12h-8a1 1 0 1 0-1 1v8a9 9 0 0 0 9-9"/><path d="M16 9a5 5 0 1 0-7 7"/><path d="M20.486 9A9 9 0 1 0 9.004 20.495"/></g>`),
		g.Group(children),
	)
}
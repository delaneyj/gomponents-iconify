package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M6 10a7 7 0 1 1 13 3.6a10 10 0 0 1-2 2a8 8 0 0 0-2 3A4.5 4.5 0 0 1 8.2 20"/><path d="M10 10a3 3 0 1 1 5 2.2"/></g>`),
		g.Group(children),
	)
}
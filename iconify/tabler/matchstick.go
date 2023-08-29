package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Matchstick(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m3 21l14-9m-1 0a1 1 0 1 0 2 0a1 1 0 1 0-2 0"/><path d="m17 3l3.62 7.29a4.007 4.007 0 0 1-.764 4.51a4 4 0 0 1-6.493-4.464L17 3z"/></g>`),
		g.Group(children),
	)
}
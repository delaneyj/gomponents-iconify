package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DropCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M10.07 15.34c1.115.88 2.74.88 3.855 0c1.115-.88 1.398-2.388.671-3.575L12 8l-2.602 3.765c-.726 1.187-.443 2.694.672 3.575z"/><path d="M3 12a9 9 0 1 0 18 0a9 9 0 1 0-18 0"/></g>`),
		g.Group(children),
	)
}
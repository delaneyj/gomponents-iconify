package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shuffle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M22 7c-3 0-8.5 0-10.5 5.5S5 18 2 18"/><path d="m20 5l2 2l-2 2m2 9c-3 0-8.5 0-10.5-5.5S5 7 2 7"/><path d="m20 20l2-2l-2-2"/></g>`),
		g.Group(children),
	)
}
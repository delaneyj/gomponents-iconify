package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandUnity(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m14 3l6 4v7m-2 3l-6 4l-6-4m-2-3V7l6-4"/><path d="m4 7l8 5v9m8-14l-8 5"/></g>`),
		g.Group(children),
	)
}
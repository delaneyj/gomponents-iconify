package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sticker(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m20 12l-2 .5A6 6 0 0 1 11.5 6l.5-2l8 8"/><path d="M20 12a8 8 0 1 1-8-8"/></g>`),
		g.Group(children),
	)
}
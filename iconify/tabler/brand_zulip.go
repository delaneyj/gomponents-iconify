package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandZulip(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M6.5 3h11C18.825 3 20 4 20 5.5c0 2-1.705 3.264-2 3.5l-4.5 4l2-5h-9a2.5 2.5 0 0 1 0-5z"/><path d="M17.5 21h-11C5.175 21 4 20 4 18.5c0-2 1.705-3.264 2-3.5l4.5-4l-2 5h9a2.5 2.5 0 1 1 0 5z"/></g>`),
		g.Group(children),
	)
}
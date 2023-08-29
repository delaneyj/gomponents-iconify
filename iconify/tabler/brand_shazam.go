package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandShazam(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m10 12l2-2a2.828 2.828 0 0 1 4 0a2.828 2.828 0 0 1 0 4l-3 3"/><path d="m14 12l-2 2a2.828 2.828 0 1 1-4-4l3-3"/><path d="M3 12a9 9 0 1 0 18 0a9 9 0 1 0-18 0"/></g>`),
		g.Group(children),
	)
}
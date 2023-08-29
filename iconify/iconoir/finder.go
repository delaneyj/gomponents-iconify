package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Finder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M3 15V9a6 6 0 0 1 6-6h6a6 6 0 0 1 6 6v6a6 6 0 0 1-6 6H9a6 6 0 0 1-6-6Z"/><path d="M15 3s-4.5 0-4.5 9H13c0 9 2 9 2 9"/><path d="M16.5 14.5s-1.5 2-4.5 2s-4.5-2-4.5-2M7 9v2m10-2v2"/></g>`),
		g.Group(children),
	)
}
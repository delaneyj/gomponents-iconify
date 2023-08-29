package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bookmarks(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M13 7a2 2 0 0 1 2 2v12l-5-3l-5 3V9a2 2 0 0 1 2-2h6z"/><path d="M9.265 4A2 2 0 0 1 11 3h6a2 2 0 0 1 2 2v12l-1-.6"/></g>`),
		g.Group(children),
	)
}
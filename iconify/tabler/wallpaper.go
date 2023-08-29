package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wallpaper(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M8 6h10a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H6"/><path d="M4 18a2 2 0 1 0 4 0a2 2 0 1 0-4 0"/><path d="M8 18V6a2 2 0 1 0-4 0v12"/></g>`),
		g.Group(children),
	)
}
package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandGit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M15 12a1 1 0 1 0 2 0a1 1 0 1 0-2 0m-4-4a1 1 0 1 0 2 0a1 1 0 1 0-2 0m0 8a1 1 0 1 0 2 0a1 1 0 1 0-2 0m1-1V9m3 2l-2-2m-2-2L9.1 5.1"/><path d="m13.446 2.6l7.955 7.954a2.045 2.045 0 0 1 0 2.892l-7.955 7.955a2.045 2.045 0 0 1-2.892 0l-7.955-7.955a2.045 2.045 0 0 1 0-2.892l7.955-7.955a2.045 2.045 0 0 1 2.892 0z"/></g>`),
		g.Group(children),
	)
}
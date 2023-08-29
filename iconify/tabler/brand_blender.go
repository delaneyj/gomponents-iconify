package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandBlender(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M9 14a6 5 0 1 0 12 0a6 5 0 1 0-12 0"/><path d="M14 14a1 1 0 1 0 2 0a1 1 0 1 0-2 0M3 16l9-6.5M6 9h9m-2-4l5.65 5"/></g>`),
		g.Group(children),
	)
}
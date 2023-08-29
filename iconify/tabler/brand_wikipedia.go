package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandWikipedia(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 4.984h2m3 0h2.5m4 0H17m5 0h-2m-16 0L9.455 19.5L16 4.984"/><path d="M9 4.984L15 19.5l6-14.516"/></g>`),
		g.Group(children),
	)
}
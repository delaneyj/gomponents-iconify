package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandFiverr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 3h-2a6 6 0 0 0-6 6H4v4h3v8h4v-7h4v7h4V10h-8V8.967A1.967 1.967 0 0 1 13 7h2V3z"/>`),
		g.Group(children),
	)
}
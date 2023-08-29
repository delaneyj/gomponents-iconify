package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SitemapF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M2 12.858h4a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2H2a2 2 0 0 1-2-2v-4a2 2 0 0 1 2-2zm6-12h4a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2H8a2 2 0 0 1-2-2v-4a2 2 0 0 1 2-2zm6 12h4a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2h-4a2 2 0 0 1-2-2v-4a2 2 0 0 1 2-2z"/><path d="M9 9.858v-2h2v4H6a1 1 0 0 0-1 1v.935H3v-.935a3 3 0 0 1 3-3h3z"/><path d="M10 11.858v-2h4a3 3 0 0 1 3 3v1.02h-2v-1.02a1 1 0 0 0-1-1h-4z"/></g>`),
		g.Group(children),
	)
}
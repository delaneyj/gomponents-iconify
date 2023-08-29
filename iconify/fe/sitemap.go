package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sitemap(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feSitemap0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feSitemap1" fill="currentColor" fill-rule="nonzero"><path id="feSitemap2" d="M19 15a2 2 0 0 1 2 2v2a2 2 0 0 1-2 2h-2a2 2 0 0 1-2-2v-2a2 2 0 0 1 2-2v-2H7v2a2 2 0 0 1 2 2v2a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-2a2 2 0 0 1 2-2v-2a2 2 0 0 1 2-2h4V9a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h2a2 2 0 0 1 2 2v2a2 2 0 0 1-2 2v2h4a2 2 0 0 1 2 2v2ZM5 17v2h2v-2H5Zm12 0v2h2v-2h-2ZM11 5v2h2V5h-2Z"/></g></g>`),
		g.Group(children),
	)
}
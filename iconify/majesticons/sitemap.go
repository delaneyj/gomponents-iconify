package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sitemap(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path fill="currentColor" d="M10 4h4v4h-4zm0 12h4v4h-4zm-8 0h4v4H2zm16 0h4v4h-4z"/><path d="M12 8v4m0 4v-4m0 0h6a2 2 0 0 1 2 2v2m-8-4H6a2 2 0 0 0-2 2v2"/></g>`),
		g.Group(children),
	)
}
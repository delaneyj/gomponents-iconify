package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandDigg(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M6 15H3v-4h3m9 4h-3v-4h3m-6 4v-4"/><path d="M15 11v7h-3M6 7v8m15 0h-3v-4h3"/><path d="M21 11v7h-3"/></g>`),
		g.Group(children),
	)
}
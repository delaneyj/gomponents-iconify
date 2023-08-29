package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandVsco(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M21 12a9 9 0 1 1-18 0a9 9 0 0 1 18 0z"/><path d="M17 12a5 5 0 1 0-10 0a5 5 0 0 0 10 0zm-5-9v4m9 5h-4m-5 9v-4m-9-5h4m11.364-6.364l-2.828 2.828m2.828 9.9l-2.828-2.828m-9.9 2.828l2.828-2.828m-2.828-9.9l2.828 2.828"/></g>`),
		g.Group(children),
	)
}
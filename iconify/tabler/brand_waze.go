package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandWaze(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M6.66 17.52A7 7 0 0 1 3 13c2 0 3-1 3-2.51C6 6.57 8.25 3 13.38 3C18 3 21 6.51 21 11a8.08 8.08 0 0 1-3.39 6.62M10 18.69a17.29 17.29 0 0 0 3.33.3h.54"/><path d="M14 19a2 2 0 1 0 4 0a2 2 0 1 0-4 0m-8 0a2 2 0 1 0 4 0a2 2 0 1 0-4 0M16 9h.01M11 9h.01"/></g>`),
		g.Group(children),
	)
}
package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandChrome(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 12a9 9 0 1 0 18 0a9 9 0 1 0-18 0"/><path d="M9 12a3 3 0 1 0 6 0a3 3 0 1 0-6 0m3-3h8.4m-5.802 4.5l-4.2 7.275M9.402 13.5l-4.2-7.275"/></g>`),
		g.Group(children),
	)
}
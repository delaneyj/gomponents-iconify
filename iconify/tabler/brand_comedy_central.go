package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandComedyCentral(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M5.343 17.657a8 8 0 1 0 0-11.314"/><path d="M13.828 9.172a4 4 0 1 0 0 5.656"/></g>`),
		g.Group(children),
	)
}
package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 18v-8l-2 2M4 6v12m8-12v12m-1 0h2M3 18h2m-1-6h8M3 6h2m6 0h2"/>`),
		g.Group(children),
	)
}
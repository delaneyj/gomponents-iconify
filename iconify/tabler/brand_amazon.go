package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandAmazon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 12.5a15.198 15.198 0 0 1-7.37 1.44A14.62 14.62 0 0 1 3 11m16.5 4c.907-1.411 1.451-3.323 1.5-5c-1.197-.773-2.577-.935-4-1"/>`),
		g.Group(children),
	)
}
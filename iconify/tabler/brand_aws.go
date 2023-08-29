package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandAws(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 18.5a15.198 15.198 0 0 1-7.37 1.44A14.62 14.62 0 0 1 3 17m16.5 4c.907-1.411 1.451-3.323 1.5-5c-1.197-.773-2.577-.935-4-1M3 11V6.5a1.5 1.5 0 0 1 3 0V11M3 9h3m3-4l1.2 6L12 7l1.8 4L15 5m3 5.25c0 .414.336.75.75.75H20a1 1 0 0 0 1-1V9a1 1 0 0 0-1-1h-1a1 1 0 0 1-1-1V6a1 1 0 0 1 1-1h1.25a.75.75 0 0 1 .75.75"/>`),
		g.Group(children),
	)
}
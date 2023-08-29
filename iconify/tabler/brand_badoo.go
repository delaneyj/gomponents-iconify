package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandBadoo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M22 9.43C22 15.268 17.523 20 12 20S2 15.338 2 9.5c0-2.667 1.83-5.01 4.322-5.429C8.814 3.653 11.222 5.463 12 8c.768-2.54 3.177-4.354 5.668-3.931C20.163 4.486 22 6.759 22 9.429z"/><path d="M7.5 10c0 2.761 2.015 5 4.5 5s4.5-2.239 4.5-5"/></g>`),
		g.Group(children),
	)
}
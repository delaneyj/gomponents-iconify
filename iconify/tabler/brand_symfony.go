package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandSymfony(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M6 13c.458.667 1.125 1 2 1c1.313 0 2-.875 2-1.5c0-1.5-2-1-2-2C8 9.875 8.516 9 9.5 9c2.5 0 1.563 2 5.5 2c.667 0 1-.333 1-1"/><path d="M9 17c-.095.667.238 1 1 1c1.714 0 2.714-2 3-6c.286-4 1.571-6 3-6c.571 0 .905.333 1 1"/><path d="M22 12c0 5.523-4.477 10-10 10S2 17.523 2 12S6.477 2 12 2a10 10 0 0 1 10 10z"/></g>`),
		g.Group(children),
	)
}
package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandTidal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m5.333 6l3.334 3.25L12 6l3.333 3.25L18.667 6L22 9.25l-3.333 3.25l-3.334-3.25L12 12.5l3.333 3.25L12 19l-3.333-3.25L12 12.5L8.667 9.25L5.333 12.5L2 9.25z"/>`),
		g.Group(children),
	)
}
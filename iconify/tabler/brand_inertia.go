package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandInertia(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m12.5 8l4 4l-4 4H17l4-4l-4-4zm-9 0l4 4l-4 4H8l4-4l-4-4z"/>`),
		g.Group(children),
	)
}
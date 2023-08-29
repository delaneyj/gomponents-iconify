package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandLinqpad(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 21h3.5l2.5-6l2.5-1l2.5 7h4l1-4.5l-2-1l-7-12L6 3l1.5 4l2.5.5l1 2.5l-7 8z"/>`),
		g.Group(children),
	)
}
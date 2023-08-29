package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PentagonTopLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M14.33 15.16L11 19.392l-5-8.66L9.33 6.5l5.33-.768l5 8.66l-5.33.768Zm2.121-2.326l-3.198.46l-1.998 2.54l-2.846-4.93l1.998-2.539l3.198-.46l2.846 4.929Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
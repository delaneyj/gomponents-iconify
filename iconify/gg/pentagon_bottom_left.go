package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PentagonBottomLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M13.33 9.232L10 5l-5 8.66l3.33 4.232l5.33.768l5-8.66l-5.33-.768Zm2.121 2.326l-3.198-.46l-1.998-2.54l-2.846 4.93l1.998 2.539l3.198.46l2.846-4.929Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PentagonTopRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M10.33 15.16L5 14.392l5-8.66l5.33.768l3.33 4.232l-5 8.66l-3.33-4.232Zm3.075.674l-1.998-2.54l-3.198-.46l2.846-4.93l3.198.461l1.998 2.54l-2.846 4.929Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ZodiacAquarius(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m15 12.41l-3-3l-3 3l-3-3l-2.29 2.3l-1.42-1.42L6 6.59l3 3l3-3l3 3l3-3l3.71 3.7l-1.42 1.42L18 9.41l-3 3m3 3l2.29 2.3l1.42-1.42l-3.71-3.7l-3 3l-3-3l-3 3l-3-3l-3.71 3.7l1.42 1.42L6 15.41l3 3l3-3l3 3l3-3Z"/>`),
		g.Group(children),
	)
}
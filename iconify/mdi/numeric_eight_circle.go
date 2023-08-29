package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumericEightCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 13h2v2h-2v-2m0-4h2v2h-2V9m0 8h2a2 2 0 0 0 2-2v-1.5a1.5 1.5 0 0 0-1.5-1.5a1.5 1.5 0 0 0 1.5-1.5V9a2 2 0 0 0-2-2h-2a2 2 0 0 0-2 2v1.5a1.5 1.5 0 0 0 1.5 1.5A1.5 1.5 0 0 0 9 13.5V15a2 2 0 0 0 2 2m1-15a10 10 0 0 1 10 10a10 10 0 0 1-10 10A10 10 0 0 1 2 12A10 10 0 0 1 12 2Z"/>`),
		g.Group(children),
	)
}
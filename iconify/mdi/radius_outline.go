package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RadiusOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2A10 10 0 0 0 2 12a10 10 0 0 0 10 10a10 10 0 0 0 10-10A10 10 0 0 0 12 2m0 18a8 8 0 0 1-8-8a8 8 0 0 1 8-8a8 8 0 0 1 8 8a8 8 0 0 1-8 8m4-5v-2h-2.28c-.36.62-1.01 1-1.72 1a2 2 0 0 1-2-2a2 2 0 0 1 2-2c.71 0 1.36.38 1.72 1H16V9l3 3l-3 3Z"/>`),
		g.Group(children),
	)
}
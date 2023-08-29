package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUpDropCircleOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22a10 10 0 0 0 10-10A10 10 0 0 0 12 2A10 10 0 0 0 2 12a10 10 0 0 0 10 10m0-2a8 8 0 0 1-8-8a8 8 0 0 1 8-8a8 8 0 0 1 8 8a8 8 0 0 1-8 8m5-6l-5-5l-5 5h10Z"/>`),
		g.Group(children),
	)
}
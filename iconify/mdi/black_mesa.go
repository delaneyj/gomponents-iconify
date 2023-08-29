package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BlackMesa(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2a10 10 0 0 1 10 10a10 10 0 0 1-10 10A10 10 0 0 1 2 12A10 10 0 0 1 12 2m0 2a8 8 0 0 0-8 8a7.99 7.99 0 0 0 2.71 6H9v-6h8l2.15 3.59c.54-1.09.85-2.3.85-3.59a8 8 0 0 0-8-8Z"/>`),
		g.Group(children),
	)
}
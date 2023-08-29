package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowLeftCircleOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 11v2h-8l3.5 3.5l-1.42 1.42L6.16 12l5.92-5.92L13.5 7.5L10 11h8M2 12A10 10 0 0 1 12 2a10 10 0 0 1 10 10a10 10 0 0 1-10 10A10 10 0 0 1 2 12m2 0a8 8 0 0 0 8 8a8 8 0 0 0 8-8a8 8 0 0 0-8-8a8 8 0 0 0-8 8Z"/>`),
		g.Group(children),
	)
}
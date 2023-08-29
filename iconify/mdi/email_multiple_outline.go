package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EmailMultipleOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 6v14h18v2H2c-1.105 0-2-.89-2-2V6h2m22-2c0-1.1-.9-2-2-2H6c-1.1 0-2 .9-2 2v12c0 1.1.9 2 2 2h16c1.1 0 2-.9 2-2V4m-2 0l-8 5l-8-5h16m0 12H6V6l8 5l8-5v10Z"/>`),
		g.Group(children),
	)
}
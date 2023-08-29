package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BeakerCheckOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m17.75 21.16l-2.75-3L16.16 17l1.59 1.59L21.34 15l1.16 1.41l-4.75 4.75M3 3h18v2c-1.1 0-2 .9-2 2v5c-.7 0-1.37.12-2 .34V5H7v2h5v1H7v1h3v1H7v1h3v1H7v1h5v1H7v1h3v1H7v3h6.08c.12.72.37 1.39.72 2H7a2 2 0 0 1-2-2V7a2 2 0 0 0-2-2V3Z"/>`),
		g.Group(children),
	)
}
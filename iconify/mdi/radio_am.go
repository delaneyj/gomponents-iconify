package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RadioAm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 7a2 2 0 0 0-2 2v8h2v-4h2v4h2V9a2 2 0 0 0-2-2H5m0 2h2v2H5V9m8-2a2 2 0 0 0-2 2v8h2V9h2v7h2V9h2v8h2V9a2 2 0 0 0-2-2h-6Z"/>`),
		g.Group(children),
	)
}
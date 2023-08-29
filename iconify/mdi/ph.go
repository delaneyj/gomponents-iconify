package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ph(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 7v10h2v-4h2v4h2V7h-2v4h-2V7h-2m-2 8v-2a2 2 0 0 0-2-2H5v10h2v-4h2c1.11 0 2-.89 2-2m-2 0H7v-2h2v2Z"/>`),
		g.Group(children),
	)
}
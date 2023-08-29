package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Castle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 13h2v2h2v-2h2v2h2v-2h2v2h2v-5l3-3V1h2l4 2l-4 2v2l3 3v12H11v-3a2 2 0 0 0-2-2a2 2 0 0 0-2 2v3H2v-9m16-3c-.55 0-1 .54-1 1.2V13h2v-1.8c0-.66-.45-1.2-1-1.2Z"/>`),
		g.Group(children),
	)
}
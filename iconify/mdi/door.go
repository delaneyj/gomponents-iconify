package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Door(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 3c-1.11 0-2 .89-2 2v16h12V5c0-1.11-.89-2-2-2H8m0 2h8v14H8V5m5 6v2h2v-2h-2Z"/>`),
		g.Group(children),
	)
}
package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PostLamp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m14 3l-1-2h-2l-1 2l-5 3h2l1 8l2 2l.5 1H9v6h6v-6h-1.5l.5-1l2-2l1-8h2l-5-3m.16 10H9.84L9 6h6l-.84 7Z"/>`),
		g.Group(children),
	)
}
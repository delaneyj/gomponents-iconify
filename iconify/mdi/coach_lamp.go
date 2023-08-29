package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CoachLamp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m16 5l-1-3h-2l-1 3l-6 3h2l.6 3H4V7H2v10h2v-4h5l1 5l2 2l1 2h2l1-2l2-2l2-10h2m-5.84 9h-4.32L10 8h8Z"/>`),
		g.Group(children),
	)
}
package bxl

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vuejs(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 12.765l5.592-9.437h-3.276L12 7.33v.002L9.688 3.328h-3.28z"/><path fill="currentColor" d="M18.461 3.332L12 14.235L5.539 3.332H1.992L12 20.672l10.008-17.34z"/>`),
		g.Group(children),
	)
}
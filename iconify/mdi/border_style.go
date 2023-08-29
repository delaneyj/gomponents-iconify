package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderStyle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 21h2v-2h-2m4 2h2v-2h-2M7 21h2v-2H7m4 2h2v-2h-2m8-2h2v-2h-2m0-2h2v-2h-2M3 3v18h2V5h16V3m-2 6h2V7h-2"/>`),
		g.Group(children),
	)
}
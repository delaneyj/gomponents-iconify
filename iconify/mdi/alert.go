package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Alert(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 14h-2V9h2m0 9h-2v-2h2M1 21h22L12 2L1 21Z"/>`),
		g.Group(children),
	)
}
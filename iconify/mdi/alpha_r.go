package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlphaR(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 7v10h2v-4h.8l1.2 4h2l-1.24-4.15C14.5 12.55 15 11.84 15 11V9a2 2 0 0 0-2-2H9m2 2h2v2h-2V9Z"/>`),
		g.Group(children),
	)
}
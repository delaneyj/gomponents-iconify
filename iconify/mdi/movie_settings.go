package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MovieSettings(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 4v14c0 1.11-.89 2-2 2H4a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h1l2 4h3L8 4h2l2 4h3l-2-4h2l2 4h3l-2-4h4M7 24h2v-2H7v2m4 0h2v-2h-2v2m4 0h2v-2h-2v2Z"/>`),
		g.Group(children),
	)
}
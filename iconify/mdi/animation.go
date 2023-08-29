package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Animation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 2c-1.11 0-2 .89-2 2v10h2V4h10V2H4m4 4c-1.11 0-2 .89-2 2v10h2V8h10V6H8m4 4c-1.11 0-2 .89-2 2v8c0 1.11.89 2 2 2h8c1.11 0 2-.89 2-2v-8c0-1.11-.89-2-2-2h-8Z"/>`),
		g.Group(children),
	)
}
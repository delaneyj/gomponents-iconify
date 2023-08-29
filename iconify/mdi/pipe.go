package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pipe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 14h-2v2h-6v-3h2v-2h-2V6a2 2 0 0 0-2-2H4V2H2v8h2V8h6v3H8v2h2v5a2 2 0 0 0 2 2h8v2h2"/>`),
		g.Group(children),
	)
}
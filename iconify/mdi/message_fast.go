package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MessageFast(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 7c-.6 0-1 .4-1 1s.4 1 1 1h2V7H3m-1 4c-.6 0-1 .4-1 1s.4 1 1 1h3v-2H2m-1 4c-.6 0-1 .4-1 1s.4 1 1 1h4v-2H1M20 5H9c-1.1 0-2 .9-2 2v14l4-4h9c1.1 0 2-.9 2-2V7c0-1.1-.9-2-2-2Z"/>`),
		g.Group(children),
	)
}
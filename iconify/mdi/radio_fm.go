package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RadioFm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 7v10h2v-4h3v-2H5V9h4V7H3m10 0a2 2 0 0 0-2 2v8h2V9h2v7h2V9h2v8h2V9a2 2 0 0 0-2-2h-6Z"/>`),
		g.Group(children),
	)
}
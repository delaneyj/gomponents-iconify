package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CupboardOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 2c-1.1 0-2 .9-2 2v15a2 2 0 0 0 2 2v1h2v-1h6v1h2v-1c1.11 0 2-.89 2-2V4a2 2 0 0 0-2-2H7m0 2h10v3H7V4m0 5h10v3H7V9m0 5h4v5H7v-5m6 0h4v5h-4v-5m-5 1v3h2v-3H8m6 0v3h2v-3h-2Z"/>`),
		g.Group(children),
	)
}
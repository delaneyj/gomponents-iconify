package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CashUsdOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M20 18H4V6h16m0-2H4c-1.11 0-2 .89-2 2v12a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2m-9 13h2v-1h1a1 1 0 0 0 1-1v-3a1 1 0 0 0-1-1h-3v-1h4V8h-2V7h-2v1h-1a1 1 0 0 0-1 1v3a1 1 0 0 0 1 1h3v1H9v2h2v1z" fill="currentColor"/>`),
		g.Group(children),
	)
}
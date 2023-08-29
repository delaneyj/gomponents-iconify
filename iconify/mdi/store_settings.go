package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StoreSettings(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 6H4V4h16v2m1 6v2h-1v6h-2v-6h-4v6H4v-6H3v-2l1-5h16l1 5m-9 2H6v4h6v-4M7 24h2v-2H7v2m4 0h2v-2h-2v2m4 0h2v-2h-2v2Z"/>`),
		g.Group(children),
	)
}
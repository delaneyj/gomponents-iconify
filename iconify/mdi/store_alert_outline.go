package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StoreAlertOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m2 7l-1 5v2h1v6h10v-6h4v6h2v-6h1v-2l-1-5H2m8 11H4v-4h6v4m-6.96-6l.6-3h12.72l.6 3H3.04M18 6H2V4h16v2m5 1v6h-2V7h2m-2 8h2v2h-2v-2Z"/>`),
		g.Group(children),
	)
}
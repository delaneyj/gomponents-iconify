package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReloadAlert(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 12c0 5 4 9 9 9c2.4 0 4.7-.9 6.4-2.6l-1.5-1.5c-1.3 1.4-3 2.1-4.9 2.1c-6.2 0-9.4-7.5-4.9-11.9S18 5.8 18 12h-3l4 4h.1l3.9-4h-3c0-5-4-9-9-9s-9 4-9 9m8 3h2v2h-2v-2m0-8h2v6h-2V7"/>`),
		g.Group(children),
	)
}
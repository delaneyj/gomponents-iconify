package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShieldSword(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 1L3 5v6c0 5.5 3.8 10.7 9 12c5.2-1.3 9-6.5 9-12V5l-9-4m3 14h-2v3h-2v-3H9v-2h2l-1-5.9l2-1.6l2 1.6l-1 5.9h2v2Z"/>`),
		g.Group(children),
	)
}
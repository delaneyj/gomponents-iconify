package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CarLimousine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m1 6l1.5 1.5L1 9l1.5 1.5L1 12l1.5 1.5L1 15h14a3 3 0 0 0 3 3a3 3 0 0 0 3-3h2v-3c0-1.11-.89-2-2-2h-2l-3-4H1m3.62 1.5h5.88V10H4.12l-1-1l1.5-1.5m7.38 0h3.5l1.96 2.5H12V7.5m6 6a1.5 1.5 0 0 1 1.5 1.5a1.5 1.5 0 0 1-1.5 1.5a1.5 1.5 0 0 1-1.5-1.5a1.5 1.5 0 0 1 1.5-1.5Z"/>`),
		g.Group(children),
	)
}
package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeatWave(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m8.5 4.5l-3.1 5l3.1 5.2l-3.3 5.8l-1.8-.9l2.7-4.9L3 9.5l3.7-5.9l1.8.9m6.2-.1l-3.1 5.1l3.1 5l-3.3 5.8l-1.8-.9l2.7-4.9l-3.1-5l3.7-6l1.8.9m6.3 0l-3.1 5.1l3.1 5l-3.3 5.8l-1.8-.9l2.7-4.9l-3.1-5l3.7-6l1.8.9"/>`),
		g.Group(children),
	)
}
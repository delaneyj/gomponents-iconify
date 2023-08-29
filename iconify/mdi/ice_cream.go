package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IceCream(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17.5 6.05C17.25 3.22 14.89 1 12 1S6.75 3.22 6.5 6.05C5.09 6.28 4 7.5 4 9c0 1.66 1.34 3 3 3l5 10l5-10c1.66 0 3-1.34 3-3c0-1.5-1.09-2.72-2.5-2.95M12 17.53l-3.11-6.22c.06-.05.11-.1.19-.15a5.503 5.503 0 0 0 5.84 0c.08.05.13.1.19.15L12 17.53Z"/>`),
		g.Group(children),
	)
}
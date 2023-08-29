package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FoodKosher(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 7v2h2c.57 0 1 .43 1 1v7h2v-7c0-1.64-1.36-3-3-3H1m6 0v7c0 1.65 1.36 3 3 3h4c1.65 0 3-1.35 3-3V7h-2v7c0 .57-.43 1-1 1h-4c-.57 0-1-.43-1-1v-1h1c1.65 0 3-1.35 3-3V7h-2v3c0 .57-.43 1-1 1H9V7H7m11 0v2h2c.57 0 1 .43 1 1v4c0 .57-.43 1-1 1h-2v2h2c1.65 0 3-1.35 3-3v-4c0-1.64-1.35-3-3-3h-2Z"/>`),
		g.Group(children),
	)
}
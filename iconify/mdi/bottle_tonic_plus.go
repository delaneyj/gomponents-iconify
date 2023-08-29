package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BottleTonicPlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 4h-2l-1-2h4l-1 2m1 4V6h1V5H9v1h1v2c-2.76 0-5 2.24-5 5v9h14v-9c0-2.76-2.24-5-5-5m2 9h-3v3h-2v-3H8v-2h3v-3h2v3h3v2Z"/>`),
		g.Group(children),
	)
}
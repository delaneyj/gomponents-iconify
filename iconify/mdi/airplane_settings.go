package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AirplaneSettings(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.09 4.56L16.2 8.45l2.12 9.19l-1.41 1.42L13 11.63L9.13 15.5L9.5 18l-1.08 1.06l-1.77-3.19l-3.18-1.76l1.06-1.07L7 13.4l3.9-3.9l-7.43-3.88l1.42-1.41l9.19 2.12l3.89-3.89a1.49 1.49 0 0 1 2.12 0c.59.56.59 1.53 0 2.12M7 24h2v-2H7v2m4 0h2v-2h-2v2m4 0h2v-2h-2v2Z"/>`),
		g.Group(children),
	)
}
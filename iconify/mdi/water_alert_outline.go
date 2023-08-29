package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WaterAlertOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 3.25S4 10 4 14c0 3.31 2.69 6 6 6s6-2.69 6-6c0-4-6-10.75-6-10.75M10 18c-2.21 0-4-1.79-4-4c0-1.77 2-5.04 4-7.61c2 2.56 4 5.84 4 7.61c0 2.21-1.79 4-4 4M20 7v6h-2V7h2m-2 10h2v-2h-2v2Z"/>`),
		g.Group(children),
	)
}
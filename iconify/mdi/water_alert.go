package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WaterAlert(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 3.25S16 10 16 14c0 3.31-2.69 6-6 6s-6-2.69-6-6c0-4 6-10.75 6-10.75M20 7v6h-2V7h2m-2 10h2v-2h-2v2Z"/>`),
		g.Group(children),
	)
}
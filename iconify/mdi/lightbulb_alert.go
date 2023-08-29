package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LightbulbAlert(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 2C6.1 2 3 5.1 3 9c0 2.4 1.2 4.5 3 5.7V17c0 .6.4 1 1 1h6c.6 0 1-.4 1-1v-2.3c1.8-1.3 3-3.4 3-5.7c0-3.9-3.1-7-7-7M7 21c0 .6.4 1 1 1h4c.6 0 1-.4 1-1v-1H7v1m12-9V7h2v6h-2m0 4v-2h2v2h-2Z"/>`),
		g.Group(children),
	)
}
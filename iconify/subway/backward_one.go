package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BackwardOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M472.8 90L298.5 192.5v-74c0-28.4-17.6-41.2-39.2-28.5L16.2 232.9c-21.6 12.7-21.6 33.4 0 46.1l243.1 143c21.6 12.7 39.2-.2 39.2-28.5v-74L472.8 422c21.6 12.7 39.2-.2 39.2-28.5v-275c0-28.4-17.6-41.2-39.2-28.5z"/>`),
		g.Group(children),
	)
}
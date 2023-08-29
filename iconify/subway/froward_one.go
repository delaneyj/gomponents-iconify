package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FrowardOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M495.8 232.9L252.7 90c-21.6-12.7-39.2.1-39.2 28.5v74L39.2 90C17.6 77.3 0 90.1 0 118.5v275c0 28.3 17.6 41.2 39.2 28.5l174.3-102.5v74c0 28.3 17.6 41.2 39.2 28.5l243.1-143c21.6-12.6 21.6-33.4 0-46.1z"/>`),
		g.Group(children),
	)
}
package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberTwoThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M172 208a4 4 0 0 1-4 4H88a4 4 0 0 1-3.2-6.4l71.94-95.92a36 36 0 1 0-57.49-43.34A36.62 36.62 0 0 0 95 73.61a4 4 0 0 1-7.33-3.21a44.42 44.42 0 0 1 5.2-8.87a44 44 0 0 1 70.28 53L96 204h72a4 4 0 0 1 4 4Z"/>`),
		g.Group(children),
	)
}
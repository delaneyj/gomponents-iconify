package bxs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FastForwardCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2C6.485 2 2 6.485 2 12s4.485 10 10 10c5.514 0 10-4.485 10-10S17.514 2 12 2zm1 14v-4l-6 4V8l6 4V8l6 4l-6 4z"/>`),
		g.Group(children),
	)
}
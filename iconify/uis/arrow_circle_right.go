package uis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowCircleRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2C6.5 2 2 6.5 2 12s4.5 10 10 10s10-4.5 10-10S17.5 2 12 2zm3.7 10.7l-3 3c-.4.4-1 .4-1.4 0c-.4-.4-.4-1 0-1.4l1.3-1.3H9c-.6 0-1-.4-1-1s.4-1 1-1h3.6l-1.3-1.3c-.4-.4-.4-1 0-1.4c.4-.4 1-.4 1.4 0l3 3c.4.4.4 1 0 1.4z"/>`),
		g.Group(children),
	)
}
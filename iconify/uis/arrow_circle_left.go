package uis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowCircleLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2C6.5 2 2 6.5 2 12s4.5 10 10 10s10-4.5 10-10S17.5 2 12 2zm3 11h-3.6l1.3 1.3c.4.4.4 1 0 1.4c-.4.4-1 .4-1.4 0l-3-3c-.4-.4-.4-1 0-1.4l3-3c.4-.4 1-.4 1.4 0c.4.4.4 1 0 1.4L11.4 11H15c.6 0 1 .4 1 1s-.4 1-1 1z"/>`),
		g.Group(children),
	)
}
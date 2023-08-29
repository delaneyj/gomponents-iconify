package uis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClockFive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2C6.5 2 2 6.5 2 12s4.5 10 10 10s10-4.5 10-10S17.5 2 12 2zm2 13.7c-.5.3-1.1.1-1.4-.4l-1.5-2.8c-.1-.2-.1-.3-.1-.5V7c0-.6.4-1 1-1s1 .4 1 1v4.7l1.4 2.6c.2.5.1 1.1-.4 1.4z"/>`),
		g.Group(children),
	)
}
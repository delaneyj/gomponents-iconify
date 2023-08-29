package uis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClockThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2C6.5 2 2 6.5 2 12s4.5 10 10 10s10-4.5 10-10S17.5 2 12 2zm3 11h-3c-.6 0-1-.4-1-1V7c0-.6.4-1 1-1s1 .4 1 1v4h2c.6 0 1 .4 1 1s-.4 1-1 1z"/>`),
		g.Group(children),
	)
}
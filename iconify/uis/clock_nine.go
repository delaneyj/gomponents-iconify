package uis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClockNine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2C6.5 2 2 6.5 2 12s4.5 10 10 10s10-4.5 10-10S17.5 2 12 2zm1 10c0 .6-.4 1-1 1H9c-.6 0-1-.4-1-1s.4-1 1-1h2V7c0-.6.4-1 1-1s1 .4 1 1v5z"/>`),
		g.Group(children),
	)
}
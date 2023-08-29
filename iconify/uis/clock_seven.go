package uis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClockSeven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2C6.5 2 2 6.5 2 12s4.5 10 10 10s10-4.5 10-10S17.5 2 12 2zm1 10c0 .2-.1.3-.1.4l-1.5 2.8c-.3.5-.9.7-1.4.4c-.5-.3-.7-.9-.4-1.4l1.4-2.6V7c0-.6.4-1 1-1s1 .4 1 1v5z"/>`),
		g.Group(children),
	)
}
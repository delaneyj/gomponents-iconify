package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlaybackSpeedLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2"/><path stroke-dasharray="4 3" stroke-linecap="round" d="M12 22C6.477 22 2 17.523 2 12S6.477 2 12 2"/><path d="M15.414 10.941c.781.462.781 1.656 0 2.118l-4.72 2.787C9.934 16.294 9 15.71 9 14.786V9.214c0-.924.934-1.507 1.694-1.059l4.72 2.787Z"/></g>`),
		g.Group(children),
	)
}
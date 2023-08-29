package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ComputerOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-width="4" d="M10 6a2 2 0 0 1 2-2h24a2 2 0 0 1 2 2v36a2 2 0 0 1-2 2H12a2 2 0 0 1-2-2V6Z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M17 12h14"/><circle cx="17" cy="21" r="2" fill="currentColor"/><circle cx="17" cy="27" r="2" fill="currentColor"/><circle cx="17" cy="33" r="2" fill="currentColor"/><circle cx="24" cy="21" r="2" fill="currentColor"/><circle cx="24" cy="27" r="2" fill="currentColor"/><circle cx="24" cy="33" r="2" fill="currentColor"/><circle cx="31" cy="21" r="2" fill="currentColor"/><circle cx="31" cy="27" r="2" fill="currentColor"/><circle cx="31" cy="33" r="2" fill="currentColor"/></g>`),
		g.Group(children),
	)
}
package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WashingMachineOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><rect width="32" height="40" x="8" y="4" stroke="currentColor" stroke-width="4" rx="2"/><path stroke="currentColor" stroke-width="4" d="M8 12a2 2 0 0 0 2 2h28a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2H10a2 2 0 0 0-2 2v6Z"/><circle cx="14" cy="9" r="2" fill="currentColor"/><circle cx="20" cy="9" r="2" fill="currentColor"/><circle cx="24" cy="29" r="7" stroke="currentColor" stroke-width="4"/></g>`),
		g.Group(children),
	)
}
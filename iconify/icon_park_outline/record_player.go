package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RecordPlayer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><rect width="38" height="32" x="5" y="8" stroke="currentColor" stroke-width="4" rx="2"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M13 8v32"/><circle cx="28" cy="24" r="9" stroke="currentColor" stroke-width="4"/><circle cx="28" cy="24" r="3" fill="currentColor"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M5 16h8m-8 8h8m-8 8h8"/></g>`),
		g.Group(children),
	)
}
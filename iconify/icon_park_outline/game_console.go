package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GameConsole(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><rect width="28" height="40" x="10" y="4" stroke="currentColor" stroke-width="4" rx="2"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M16 34h8m-4-4v8"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="4" d="M16 10h16v9H16z"/><circle cx="31" cy="30" r="2" fill="currentColor"/><circle cx="31" cy="38" r="2" fill="currentColor"/></g>`),
		g.Group(children),
	)
}
package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GameConsoleOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><rect width="28" height="40" x="10" y="4" stroke="currentColor" stroke-width="4" rx="2"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M16 12h16v10H16zm0 20h8m-4 4v-8"/><rect width="4" height="4" x="27" y="33" fill="currentColor" rx="2"/><rect width="4" height="4" x="30" y="26" fill="currentColor" rx="2"/></g>`),
		g.Group(children),
	)
}
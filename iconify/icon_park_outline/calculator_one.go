package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalculatorOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><rect width="32" height="40" x="8" y="4" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" rx="2"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M14 11h20v9H14z"/><circle cx="17" cy="26" r="2" fill="currentColor"/><circle cx="17" cy="32" r="2" fill="currentColor"/><circle cx="17" cy="38" r="2" fill="currentColor"/><circle cx="24" cy="26" r="2" fill="currentColor"/><circle cx="24" cy="32" r="2" fill="currentColor"/><circle cx="24" cy="38" r="2" fill="currentColor"/><circle cx="31" cy="26" r="2" fill="currentColor"/><circle cx="31" cy="32" r="2" fill="currentColor"/><circle cx="31" cy="38" r="2" fill="currentColor"/></g>`),
		g.Group(children),
	)
}
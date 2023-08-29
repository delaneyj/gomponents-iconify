package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GameConsoleOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><rect width="28" height="40" x="10" y="4" stroke="#000" stroke-width="4" rx="2"/><rect width="16" height="10" x="16" y="12" fill="#2F88FF" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M16 32L24 32"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M20 36V28"/><rect width="4" height="4" x="27" y="33" fill="#000" rx="2"/><rect width="4" height="4" x="30" y="26" fill="#000" rx="2"/></g>`),
		g.Group(children),
	)
}
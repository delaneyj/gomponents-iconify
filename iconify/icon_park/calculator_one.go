package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalculatorOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><rect width="32" height="40" x="8" y="4" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" rx="2"/><rect width="20" height="9" x="14" y="11" fill="#2F88FF" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"/><circle cx="17" cy="26" r="2" fill="#000"/><circle cx="17" cy="32" r="2" fill="#000"/><circle cx="17" cy="38" r="2" fill="#000"/><circle cx="24" cy="26" r="2" fill="#000"/><circle cx="24" cy="32" r="2" fill="#000"/><circle cx="24" cy="38" r="2" fill="#000"/><circle cx="31" cy="26" r="2" fill="#000"/><circle cx="31" cy="32" r="2" fill="#000"/><circle cx="31" cy="38" r="2" fill="#000"/></g>`),
		g.Group(children),
	)
}
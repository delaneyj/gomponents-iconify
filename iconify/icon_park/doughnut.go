package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Doughnut(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="4"><circle cx="24" cy="24" r="19"/><circle cx="24" cy="24" r="7" fill="#2F88FF"/><path stroke-linecap="round" stroke-linejoin="round" d="M6 28C6 28 9.38888 25.0417 12 27C16 30 18 28 18 28"/><path stroke-linecap="round" stroke-linejoin="round" d="M29 29C29.6667 28 33 24.7143 36 26C40 27.7143 43 26 43 25"/></g>`),
		g.Group(children),
	)
}
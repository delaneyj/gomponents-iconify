package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SlippersOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M4 29H44V35H4V29Z"/><path fill="#2F88FF" d="M7.00044 22C4 26 4 29 4 29H30.0007C30.0007 29 30.0003 23.5 30.0003 21C30.0003 18.5 28.5001 15.5 25.0001 15C21.5001 14.5 11.8762 15.5 7.00044 22Z"/></g>`),
		g.Group(children),
	)
}
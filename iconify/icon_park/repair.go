package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Repair(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M17 14L29 14"/><path d="M23 28L23 15"/><path d="M20.1429 42H8C6.89543 42 6 41.1046 6 40V7C6 5.89543 6.89543 5 8 5H40C41.1046 5 42 5.89543 42 7V16.7167"/><path fill="#2F88FF" d="M27 38L37.5 23.5L42 27L31 42H27V38Z"/></g>`),
		g.Group(children),
	)
}
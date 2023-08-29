package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EggOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path d="M29 22.133C29 30.97 23.627 36 17 36S5 30.97 5 22.133C5 11.163 11.373 4 17 4s12 7.163 12 18.133Z"/><path d="M29 24.068c8.536.634 14 5.554 14 9.932c0 4.69-6.268 10-15.867 10c-6.773 0-10.99-3.436-11.932-8"/></g>`),
		g.Group(children),
	)
}
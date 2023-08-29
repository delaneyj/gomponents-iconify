package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lipstick(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="12" height="20" x="29" y="24"/><rect width="14" height="20" x="7" y="24"/><path fill="#2F88FF" d="M10 11.4545V24H18V4C11.5 4 10 9.63636 10 11.4545Z"/><path d="M7 32L21 32"/></g>`),
		g.Group(children),
	)
}
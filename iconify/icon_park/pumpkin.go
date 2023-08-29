package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pumpkin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="4"><rect width="40" height="26" x="4" y="14" fill="#2F88FF" rx="13"/><ellipse cx="24" cy="27" rx="8" ry="13"/><path stroke-linecap="round" stroke-linejoin="round" d="M30 6H27C25.3431 6 24 7.34315 24 9V14"/></g>`),
		g.Group(children),
	)
}
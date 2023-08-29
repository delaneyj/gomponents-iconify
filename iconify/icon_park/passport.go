package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Passport(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M10 10H32H38V44H10V10Z"/><path stroke="#000" d="M10 10L32 4V10"/><circle cx="24" cy="24" r="4" fill="#2F88FF" stroke="#fff"/><path stroke="#fff" d="M20 34H28"/></g>`),
		g.Group(children),
	)
}
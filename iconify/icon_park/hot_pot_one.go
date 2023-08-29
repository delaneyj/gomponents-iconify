package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HotPotOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" stroke-linejoin="round" d="M24 36C36 36 41 27.9411 41 18H7C7 27.9411 12 36 24 36Z"/><path stroke="#000" stroke-linejoin="round" d="M17 35L14 44H34L31 35"/><path stroke="#000" stroke-linejoin="round" d="M29 18L27.8889 4L20.1111 4L19 18"/><path stroke="#fff" d="M15 25C15 25 15.0703 26.0703 16 27C16.9297 27.9297 18 28 18 28"/></g>`),
		g.Group(children),
	)
}
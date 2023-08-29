package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScaleOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linejoin="round" stroke-width="4"><path stroke="#000" stroke-linecap="round" d="M18 14L5 14L5 9L18 4L30 4L43 9L43 14L30 14"/><rect width="12" height="40" x="18" y="4" fill="#2F88FF" stroke="#000"/><path stroke="#fff" stroke-linecap="round" d="M18 12H22"/><path stroke="#fff" stroke-linecap="round" d="M18 30H23"/><path stroke="#fff" stroke-linecap="round" d="M18 18H23"/><path stroke="#fff" stroke-linecap="round" d="M18 24H22"/><path stroke="#fff" stroke-linecap="round" d="M18 36H22"/><path stroke="#000" stroke-linecap="round" d="M18 10V38"/></g>`),
		g.Group(children),
	)
}
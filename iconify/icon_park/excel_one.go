package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExcelOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-width="4"><rect width="36" height="36" x="6" y="6" fill="#2F88FF" stroke="#000" rx="3"/><path stroke="#fff" stroke-linecap="round" d="M30 16H18"/><path stroke="#fff" stroke-linecap="round" d="M18 32L18 16"/><path stroke="#fff" stroke-linecap="round" d="M28 24H18"/><path stroke="#fff" stroke-linecap="round" d="M30 32H18"/></g>`),
		g.Group(children),
	)
}
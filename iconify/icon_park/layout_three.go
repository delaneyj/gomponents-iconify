package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LayoutThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linejoin="round" stroke-width="4"><rect width="36" height="36" x="6" y="6" fill="#2F88FF" stroke="#000" rx="3"/><path stroke="#fff" stroke-linecap="round" d="M6 22L42 22"/><path stroke="#fff" stroke-linecap="round" d="M29 22V6"/><path stroke="#000" stroke-linecap="round" d="M26 6H32"/><path stroke="#000" stroke-linecap="round" d="M6 19V25"/><path stroke="#000" stroke-linecap="round" d="M42 19V25"/></g>`),
		g.Group(children),
	)
}
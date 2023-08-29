package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CameraTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><rect width="24" height="32" x="12" y="4" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" rx="12"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M24 36V44"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M16 44L32 44"/><circle cx="24" cy="17" r="6" fill="#2F88FF" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"/><circle cx="24" cy="29" r="2" fill="#000"/></g>`),
		g.Group(children),
	)
}
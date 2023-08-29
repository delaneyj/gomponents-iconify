package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Refrigerator(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linejoin="round" stroke-width="4"><rect width="28" height="36" x="9" y="4" fill="#2F88FF" stroke="#000" rx="2"/><path stroke="#fff" stroke-linecap="round" d="M9 22H37"/><path stroke="#000" stroke-linecap="round" d="M9 20L9 24"/><path stroke="#000" stroke-linecap="round" d="M37 20L37 24"/><path stroke="#fff" stroke-linecap="round" d="M15 29L15 33"/><path stroke="#fff" stroke-linecap="round" d="M15 11L15 15"/><path stroke="#000" stroke-linecap="round" d="M33 40V44"/><path stroke="#000" stroke-linecap="round" d="M13 40V44"/></g>`),
		g.Group(children),
	)
}
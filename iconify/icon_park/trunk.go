package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Trunk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="24" height="28" x="12" y="12" fill="#2F88FF" stroke="#000" rx="4"/><path stroke="#000" d="M20 12V6"/><path stroke="#000" d="M28 12V6"/><path stroke="#000" d="M16 4H32"/><path stroke="#000" d="M18 40V44"/><path stroke="#000" d="M30 40V44"/><path stroke="#fff" d="M20 25H24H28"/></g>`),
		g.Group(children),
	)
}
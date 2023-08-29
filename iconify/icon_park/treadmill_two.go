package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TreadmillTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M8 39V44"/><path stroke-linecap="round" stroke-linejoin="round" d="M40 39V44"/><path stroke-linecap="round" stroke-linejoin="round" d="M36 31L42 15L38 9"/><path stroke-linecap="round" stroke-linejoin="round" d="M33 14L43 4"/><rect width="40" height="8" x="4" y="31" fill="#2F88FF" rx="4"/></g>`),
		g.Group(children),
	)
}
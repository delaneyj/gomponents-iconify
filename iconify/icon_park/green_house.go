package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GreenHouse(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><path d="M42 20V44H24H6V20L24 4L42 20Z"/><path stroke-linecap="round" d="M6 24L42 24"/><path stroke-linecap="round" d="M13 14L13 44"/><path stroke-linecap="round" d="M35 14L35 44"/><rect width="8" height="12" x="20" y="32" fill="#2F88FF" stroke-linecap="round"/></g>`),
		g.Group(children),
	)
}
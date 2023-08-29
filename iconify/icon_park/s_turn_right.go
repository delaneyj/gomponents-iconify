package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func STurnRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M30 24L13 24C9 24 6 21 6 17C6 13 9 10 13 10L32 10"/><path stroke-linecap="round" stroke-linejoin="round" d="M8 38L35 38C39 38 42 35 42 31C42 27 39 24 35 24L30 24"/><path stroke-linecap="round" stroke-linejoin="round" d="M13 43L8 38L13 33"/><circle cx="37.176" cy="10" r="5" fill="#2F88FF" transform="rotate(-180 37.176 10)"/></g>`),
		g.Group(children),
	)
}
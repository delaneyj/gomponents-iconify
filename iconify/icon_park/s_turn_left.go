package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func STurnLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M42 38L13 38C8.99999 38 5.99999 35 5.99999 31C5.99999 27 8.99999 24 13 24L19 24"/><path stroke-linecap="round" stroke-linejoin="round" d="M19 24L35 24C39 24 42 21 42 17C42 13 39 10 35 10L17 10"/><path stroke-linecap="round" stroke-linejoin="round" d="M37 43L42 38L37 33"/><circle cx="12" cy="10" r="5" fill="#2F88FF" transform="rotate(-180 12 10)"/></g>`),
		g.Group(children),
	)
}
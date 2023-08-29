package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UTurnUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M13 14L13 33C13 39.0751 17.9249 44 24 44V44C30.0751 44 35 39.0751 35 33V4"/><path stroke-linecap="round" stroke-linejoin="round" d="M30 9L35 4L40 9"/><circle cx="13" cy="9" r="5" fill="#2F88FF" transform="rotate(-90 13 9)"/></g>`),
		g.Group(children),
	)
}
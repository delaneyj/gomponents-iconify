package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SplitTurnDownRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M13 22H29C33.4183 22 37 25.5817 37 30V44"/><circle cx="13" cy="8.944" r="5" fill="#2F88FF" transform="rotate(-90 13 8.944)"/><path stroke-linecap="round" stroke-linejoin="round" d="M13 14V43"/><path stroke-linecap="round" stroke-linejoin="round" d="M18 39L13 44L8 39"/><path stroke-linecap="round" stroke-linejoin="round" d="M42 39L37 44L32 39"/></g>`),
		g.Group(children),
	)
}
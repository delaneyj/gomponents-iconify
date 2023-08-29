package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UTurnDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M35 34V15C35 8.92487 30.0751 4 24 4V4C17.9249 4 13 8.92487 13 15V44"/><path stroke-linecap="round" stroke-linejoin="round" d="M18 39L13 44L8 39"/><circle cx="35" cy="39" r="5" fill="#2F88FF" transform="rotate(90 35 39)"/></g>`),
		g.Group(children),
	)
}
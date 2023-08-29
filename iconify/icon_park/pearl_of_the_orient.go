package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PearlOfTheOrient(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="4"><circle cx="24" cy="33" r="6" fill="#2F88FF"/><circle cx="24" cy="13" r="3" fill="#2F88FF"/><path stroke-linecap="round" d="M24 40V44"/><path stroke-linecap="round" d="M22 16L20 28"/><path stroke-linecap="round" d="M26 16L28 28"/><path stroke-linecap="round" d="M19 37L16 44"/><path stroke-linecap="round" d="M29 37L32 44"/><path stroke-linecap="round" d="M24 4V10"/></g>`),
		g.Group(children),
	)
}
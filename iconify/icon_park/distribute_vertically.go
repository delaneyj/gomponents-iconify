package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DistributeVertically(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><rect width="28" height="12" x="38" y="30" fill="#2F88FF" transform="rotate(180 38 30)"/><path stroke-linecap="round" d="M42 40H6"/><path stroke-linecap="round" d="M42 8L6 8"/></g>`),
		g.Group(children),
	)
}
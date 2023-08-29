package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoreFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="#2F88FF" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="10" height="10" x="17" y="11" transform="rotate(-45 17 11)"/><rect width="10" height="10" x="30" y="24" transform="rotate(-45 30 24)"/><rect width="10" height="10" x="4" y="24" transform="rotate(-45 4 24)"/><rect width="10" height="10" x="17" y="37" transform="rotate(-45 17 37)"/></g>`),
		g.Group(children),
	)
}
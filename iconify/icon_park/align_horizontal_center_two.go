package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignHorizontalCenterTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="#2F88FF" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="6" height="24" x="7" y="12"/><rect width="6" height="32" x="21" y="8"/><rect width="6" height="18" x="35" y="15"/></g>`),
		g.Group(children),
	)
}
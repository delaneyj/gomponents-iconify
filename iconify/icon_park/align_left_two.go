package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignLeftTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="#2F88FF" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="24" height="6" x="8" y="7"/><rect width="32" height="6" x="8" y="21"/><rect width="18" height="6" x="8" y="35"/></g>`),
		g.Group(children),
	)
}
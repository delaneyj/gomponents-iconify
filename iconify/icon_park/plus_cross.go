package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlusCross(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#2F88FF" stroke="#000" stroke-linejoin="round" stroke-width="4" d="M30 4H18V18H4V30H18V44H30V30H44V18H30V4Z"/>`),
		g.Group(children),
	)
}
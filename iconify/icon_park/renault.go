package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Renault(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="#2F88FF" stroke="#000" stroke-linejoin="round" stroke-width="4"><path d="M30 4H18L8 24H17L22 14H26L31 24H40L30 4Z"/><path d="M30 44H18L8 24H17L22 34H26L31 24H40L30 44Z"/></g>`),
		g.Group(children),
	)
}
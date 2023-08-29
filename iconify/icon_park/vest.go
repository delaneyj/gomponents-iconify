package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vest(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="#2F88FF" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M36 4H31L24 24V44H40V19L36 13V4Z"/><path d="M12 4H17L24 24V44H8V19L12 13V4Z"/></g>`),
		g.Group(children),
	)
}
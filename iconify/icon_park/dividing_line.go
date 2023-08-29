package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DividingLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M5 24H43"/><path d="M21 38H27"/><path d="M37 38H43"/><path d="M21 10H27"/><path d="M5 38H11"/><path d="M5 10H11"/><path d="M37 10H43"/></g>`),
		g.Group(children),
	)
}
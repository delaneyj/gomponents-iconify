package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AddPrint(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-width="4"><path stroke="#000" stroke-linejoin="round" d="M12 19H6V6H42V19H36"/><path fill="#2F88FF" stroke="#000" stroke-linejoin="round" d="M12 12H36V44L30 39.5556L24 44L18 39.5556L12 44V12Z"/><path stroke="#fff" d="M20 26H28"/><path stroke="#fff" d="M24 22L24 30"/></g>`),
		g.Group(children),
	)
}
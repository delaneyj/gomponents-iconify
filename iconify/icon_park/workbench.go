package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Workbench(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M12 33H4V7H44V33H36H12Z"/><path stroke="#fff" stroke-linecap="round" d="M16 22V26"/><path stroke="#000" stroke-linecap="round" d="M24 33V39"/><path stroke="#fff" stroke-linecap="round" d="M24 18V26"/><path stroke="#fff" stroke-linecap="round" d="M32 14V26"/><path stroke="#000" stroke-linecap="round" d="M12 41H36"/></g>`),
		g.Group(children),
	)
}
package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SecurityStall(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-width="4"><rect width="16" height="15" x="16" y="29" fill="#2F88FF" stroke-linejoin="round"/><rect width="24" height="6" x="12" y="4" fill="#2F88FF" stroke-linejoin="round"/><path d="M16 10V29"/><path d="M32 10V29"/><path stroke-linejoin="round" d="M4 44H44"/></g>`),
		g.Group(children),
	)
}
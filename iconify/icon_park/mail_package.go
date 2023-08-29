package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailPackage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><path d="M4 42H44V30V18H24H4V30V42Z"/><path stroke-linecap="round" d="M4 18L24 33L44 18"/><path stroke-linecap="round" d="M24 18H4V33"/><path stroke-linecap="round" d="M44 33V18H24"/><path stroke-linecap="round" d="M12 12H36"/><path stroke-linecap="round" d="M16 6H32"/></g>`),
		g.Group(children),
	)
}
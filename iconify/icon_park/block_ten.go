package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BlockTen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="#2F88FF" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M24 6H36V18H24V6Z"/><path d="M24 6H36V18H24V6Z"/><path d="M12 6H24V18H12V6Z"/><path d="M12 30H24V42H12V30Z"/><path d="M12 18H24V30H12V18Z"/></g>`),
		g.Group(children),
	)
}
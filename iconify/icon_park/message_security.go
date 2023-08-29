package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MessageSecurity(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M25.5 37H21L11 42V37H4V7H44V18"/><path fill="#2F88FF" d="M29 25.2C29 24.1333 36 22 36 22C36 22 43 24.1333 43 25.2C43 33.7333 36 38 36 38C36 38 29 33.7333 29 25.2Z"/><path d="M12 15H15L18 15"/><path d="M12 21H18L24 21"/></g>`),
		g.Group(children),
	)
}
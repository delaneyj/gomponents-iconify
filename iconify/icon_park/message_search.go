package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MessageSearch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M25.5 37H21L11 42V37H4V7H44V18"/><circle cx="34" cy="28" r="6" fill="#2F88FF"/><path stroke-linecap="round" stroke-linejoin="round" d="M39 32L44 36"/><path stroke-linecap="round" stroke-linejoin="round" d="M12 15H15L18 15"/><path stroke-linecap="round" stroke-linejoin="round" d="M12 21H18L24 21"/></g>`),
		g.Group(children),
	)
}
package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MessageFailed(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M25.5 37H21L11 42V37H4V7H44V18"/><path d="M12 15H15L18 15"/><path d="M12 21H18L24 21"/><path d="M32 25L44 37"/><path d="M44 25L32 37"/></g>`),
		g.Group(children),
	)
}
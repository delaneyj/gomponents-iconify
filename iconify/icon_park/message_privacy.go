package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MessagePrivacy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M25.5 37H21L11 42V37H4V7H44V18"/><rect width="14" height="8" x="30" y="27" fill="#2F88FF"/><path d="M40 27V24C40 22.3431 38.6569 21 37 21C35.3431 21 34 22.3431 34 24V27"/><path d="M12 15H15L18 15"/><path d="M12 21H18L24 21"/></g>`),
		g.Group(children),
	)
}
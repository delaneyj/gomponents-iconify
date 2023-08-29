package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ConnectAddressOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M6 24C6 33.9411 14.0589 42 24 42C33.9411 42 42 33.9411 42 24"/><path d="M24 14L24 42"/><circle cx="24" cy="10" r="4" fill="#2F88FF"/></g>`),
		g.Group(children),
	)
}
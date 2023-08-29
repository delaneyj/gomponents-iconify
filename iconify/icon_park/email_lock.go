package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EmailLock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M44 24V9H24H4V24V39H24"/><path d="M4 9L24 24L44 9"/><rect width="12" height="8" x="31" y="33" fill="#2F88FF"/><path d="M40 33V30C40 28.3431 38.6569 27 37 27C35.3431 27 34 28.3431 34 30V33"/></g>`),
		g.Group(children),
	)
}
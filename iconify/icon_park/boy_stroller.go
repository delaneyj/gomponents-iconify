package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoyStroller(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" d="M40 24H12L14.5 34H36L40 24Z"/><path d="M12 24L8 15H3.5"/><circle cx="20" cy="41" r="3" fill="#2F88FF"/><circle cx="31" cy="41" r="3" fill="#2F88FF"/><path fill="#2F88FF" d="M23 8L32 24H40L44 13C41.6667 10 37 4 34 4C30 4 25.6667 6.66667 23 8Z"/><path d="M29 5L33 12"/></g>`),
		g.Group(children),
	)
}
package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrendTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M4 44H44"/><path fill="#2F88FF" d="M4 26L12 28V38H4V26Z"/><path fill="#2F88FF" d="M20 24L28 20V38H20V24Z"/><path fill="#2F88FF" d="M36 16L44 12V38H36V16Z"/><path stroke-linecap="round" d="M4 18L12 20L44 4H34"/></g>`),
		g.Group(children),
	)
}
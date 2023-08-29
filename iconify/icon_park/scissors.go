package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Scissors(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="38" height="38" x="5" y="5" fill="#2F88FF" stroke="#000" rx="3"/><path stroke="#fff" d="M19 19C21 22 21 26 19 29"/><path stroke="#fff" d="M36 14L21 24L36 34"/><circle cx="16" cy="16" r="4" stroke="#fff"/><circle cx="16" cy="32" r="4" stroke="#fff"/></g>`),
		g.Group(children),
	)
}
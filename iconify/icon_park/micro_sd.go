package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicroSd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M13 7C13 5.89543 13.8954 5 15 5H37C38.1046 5 39 5.89543 39 7V41C39 42.1046 38.1046 43 37 43H10C8.89543 43 8 42.1046 8 41V36L13 32V28H8V25L13 19V7Z"/><path stroke="#fff" d="M32 11V15"/><path stroke="#fff" d="M20 11V15"/><path stroke="#fff" d="M26 11V15"/></g>`),
		g.Group(children),
	)
}
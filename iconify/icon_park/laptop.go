package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Laptop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="32" height="20" x="8" y="8" fill="#2F88FF"/><path d="M8 28L4 41H44L40 28"/><path fill="#2F88FF" d="M19.9 35H28.1L30 41H18L19.9 35Z"/></g>`),
		g.Group(children),
	)
}
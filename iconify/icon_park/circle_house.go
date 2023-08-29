package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleHouse(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M44 23H4C4 23 14.5 17 19 12C23.5 7 24.5 4 24.5 4C24.5 4 25.5 7 30 12C34.5 17 44 23 44 23Z"/><rect width="32" height="13" x="8" y="31" fill="#2F88FF"/><rect width="22" height="8" x="13" y="23" fill="#2F88FF"/></g>`),
		g.Group(children),
	)
}
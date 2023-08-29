package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PeopleSafeOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><circle cx="24" cy="11" r="7" fill="#2F88FF"/><path d="M4 41C4 32.1634 12.0589 25 22 25"/><path fill="#2F88FF" d="M28 29.2C28 28.1333 35 26 35 26C35 26 42 28.1333 42 29.2C42 37.7333 35 42 35 42C35 42 28 37.7333 28 29.2Z"/></g>`),
		g.Group(children),
	)
}
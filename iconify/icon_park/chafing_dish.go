package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChafingDish(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M24 42C14.0589 42 6 33.9411 6 24C6 14.0589 14.0589 6 24 6"/><path fill="#2F88FF" d="M24 42C33.9411 42 42 33.9411 42 24C42 14.0589 33.9411 6 24 6C24 6 20 8 20 15C20 22 28 26 28 33C28 40 24 42 24 42Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M42 20H44V28H42"/><path stroke-linecap="round" stroke-linejoin="round" d="M6 20H4V28H6"/></g>`),
		g.Group(children),
	)
}
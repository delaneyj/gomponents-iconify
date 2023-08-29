package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FThreeKey(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="36" height="36" x="6" y="6" fill="#2F88FF" stroke="#000" rx="3"/><path stroke="#fff" d="M27 16H34L29.1 23.2C31.9 23.2 34 25 34 28C34 31 32 32 30.5 32C28.8333 32 27.5 31 27 30"/><path stroke="#fff" d="M21 16H14V32"/><path stroke="#fff" d="M14 24H21"/></g>`),
		g.Group(children),
	)
}
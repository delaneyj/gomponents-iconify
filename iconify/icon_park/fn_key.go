package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FNKey(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-width="4"><rect width="36" height="36" x="6" y="6" fill="#2F88FF" stroke="#000" stroke-linejoin="round" rx="3"/><path stroke="#fff" d="M26 17V32"/><path stroke="#fff" d="M26 23C26 20.6214 27.6 19 30 19C32.4 19 34 20.5273 34 23C34 24.6484 34 28.707 34 32"/><path stroke="#fff" stroke-linejoin="round" d="M21 16H14V32"/><path stroke="#fff" stroke-linejoin="round" d="M14 24H21"/></g>`),
		g.Group(children),
	)
}
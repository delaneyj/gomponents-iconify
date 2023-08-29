package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GlassesOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><circle cx="14.5" cy="24.5" r="6.5" fill="#2F88FF"/><circle r="6.5" fill="#2F88FF" transform="matrix(-1 0 0 1 33.5 24.5)"/><path d="M4 24H8"/><path d="M44 24H40"/><path d="M20 21C20.5 19 22 17 24 17C26 17 27.5 19 28 21"/></g>`),
		g.Group(children),
	)
}
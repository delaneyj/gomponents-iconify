package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MirrorTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><circle cx="24" cy="17" r="13" fill="#2F88FF"/><path d="M42 17C42 26.9411 33.9411 35 24 35C14.0589 35 6 26.9411 6 17"/><path d="M42 17H38"/><path d="M10 17H6"/><path d="M30 44H18"/><path d="M24 44V36"/></g>`),
		g.Group(children),
	)
}
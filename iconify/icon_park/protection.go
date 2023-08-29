package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Protection(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="4"><path fill="#2F88FF" stroke-linejoin="round" d="M42 24C42 14.6112 33.9411 7 24 7C14.0589 7 6 14.6112 6 24H42Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M24 24.0083V38.5454C24 41.5579 21.5142 44 18.5 44C15.4858 44 13 41.5579 13 38.5454"/><path stroke-linecap="round" d="M24 4V7"/></g>`),
		g.Group(children),
	)
}
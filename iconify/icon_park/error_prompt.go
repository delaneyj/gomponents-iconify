package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ErrorPrompt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M8 18L20 30"/><path d="M20 18L8 30"/><path d="M34 8C39.0007 12.3609 42 17.9311 42 24C42 30.0689 39.0007 35.6391 34 40"/><path d="M27 14C30.7505 16.7256 33 20.2069 33 24C33 27.7931 30.7505 31.2744 27 34"/></g>`),
		g.Group(children),
	)
}
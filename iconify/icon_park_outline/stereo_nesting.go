package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StereoNesting(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M17 40L4 33V17l13-7l13 7v12"/><path d="m30 8l13 7v16l-13 7l-13-7V19"/></g>`),
		g.Group(children),
	)
}
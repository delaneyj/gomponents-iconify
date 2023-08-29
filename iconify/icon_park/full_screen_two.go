package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FullScreenTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M30 6H42V18"/><path d="M18 6H6V18"/><path d="M30 42H42V30"/><path d="M18 42H6V30"/><path d="M42 6L29 19"/><path d="M19 29L6 42"/></g>`),
		g.Group(children),
	)
}
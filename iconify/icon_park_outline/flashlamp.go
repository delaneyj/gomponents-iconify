package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flashlamp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><circle cx="24" cy="24" r="20"/><path stroke-linecap="round" stroke-linejoin="round" d="m23 14l-5 10h12l-5 10"/></g>`),
		g.Group(children),
	)
}
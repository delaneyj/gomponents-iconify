package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AutoHeightOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="4"><path d="M7 42V6"/><path stroke-linejoin="round" d="M18 13.99L23.995 8L30 14m0 20.01L24.005 40L18 34m6-26v32"/><path d="M41 42V6"/></g>`),
		g.Group(children),
	)
}
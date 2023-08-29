package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RotationHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="m20 31l4 4l-4 4"/><path d="M32 34.168C39.064 32.625 44 29.1 44 25c0-5.523-8.954-10-20-10S4 19.477 4 25s8.954 10 20 10"/></g>`),
		g.Group(children),
	)
}
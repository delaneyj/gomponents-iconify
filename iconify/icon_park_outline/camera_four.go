package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CameraFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path d="M8 10v14c0 8.837 7.163 16 16 16s16-7.163 16-16V10"/><path stroke-linecap="round" stroke-linejoin="round" d="M4 10h40"/><path stroke-linejoin="round" d="M24 30a6 6 0 1 0 0-12a6 6 0 0 0 0 12Z"/></g>`),
		g.Group(children),
	)
}
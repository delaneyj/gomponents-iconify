package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GreenHouse(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path d="M42 20v24H6V20L24 4l18 16Z"/><path stroke-linecap="round" d="M6 24h36M13 14v30m22-30v30M20 32h8v12h-8z"/></g>`),
		g.Group(children),
	)
}
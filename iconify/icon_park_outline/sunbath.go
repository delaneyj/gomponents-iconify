package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sunbath(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="4"><path d="M11 36v4"/><path stroke-linejoin="round" d="M6 36h34M8 30h16m13-19l-3 5"/><path d="M35 36v4"/><path stroke-linejoin="round" d="M44 10L29 36m10-17l5 9"/></g>`),
		g.Group(children),
	)
}
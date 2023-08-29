package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FigmaResetInstance(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M44 24L24 4L4 24l20 20l11-11"/><path d="M22 24s12.5-.5 17.5 4.5S39 44 39 44"/><path d="m26 20l-4 4l4 4"/></g>`),
		g.Group(children),
	)
}
package humbleicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CornerRightUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path d="M5 20h6.5a3 3 0 0 0 3-3V5"/><path d="M18 7.5L14.5 4L11 7.5"/></g>`),
		g.Group(children),
	)
}
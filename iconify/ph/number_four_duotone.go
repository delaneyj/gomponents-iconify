package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberFourDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M216 40v176a16 16 0 0 1-16 16H56a16 16 0 0 1-16-16V40a16 16 0 0 1 16-16h144a16 16 0 0 1 16 16Z" opacity=".2"/><path d="M184 152a8 8 0 0 1-8 8h-16v48a8 8 0 0 1-16 0v-48H72a8 8 0 0 1-7.53-10.69l40-112a8 8 0 0 1 15.06 5.38L83.35 144H144V96a8 8 0 0 1 16 0v48h16a8 8 0 0 1 8 8Z"/></g>`),
		g.Group(children),
	)
}
package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GraphicEq(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M85 344V88h43v256H85zm86 85V3h42v426h-42zM0 259v-86h43v86H0zm256 85V88h43v256h-43zm85-171h43v86h-43v-86z"/>`),
		g.Group(children),
	)
}
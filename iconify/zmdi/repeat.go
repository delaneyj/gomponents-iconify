package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Repeat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M85 109v86H43V67h256V3l85 85l-85 85v-64H85zm214 214v-86h42v128H85v64L0 344l85-85v64h214z"/>`),
		g.Group(children),
	)
}
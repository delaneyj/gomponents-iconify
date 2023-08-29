package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowReturnLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.935 20.822H24V7.414H5.746V3.2L0 8.946l5.746 5.746v-4.213h15.189v7.28H2.988v3.065z"/>`),
		g.Group(children),
	)
}
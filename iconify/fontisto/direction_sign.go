package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DirectionSign(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.497 17.679h4.437V24H9.497zM1.906 4.054L-.001 5.961l1.907 1.906h12.026V4.053zm0 9.331l-1.907 1.906l1.907 1.9h12.026v-3.806zm19.616-4.662H9.736v3.814h11.786l1.907-1.907zM11.71 0L9.496 2.545v1.028h4.437V2.545z"/>`),
		g.Group(children),
	)
}
package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AngleDobuleUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m11.613 10.065l11.613 11.591L20.874 24l-9.261-9.239L2.352 24L0 21.656zm0-10.065l11.613 11.591l-2.352 2.344l-9.261-9.239l-9.261 9.239L0 11.591z"/>`),
		g.Group(children),
	)
}
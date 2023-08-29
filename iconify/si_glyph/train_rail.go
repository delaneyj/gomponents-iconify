package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrainRail(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M4.989 16h-.978l2-16h.978l-2 16zm9 0h-.978l-2-16h.978l2 16z"/><path d="M3 13h12v.916H3zm0-3h12v.916H3zm1-3h10v.916H4zm0-3h10v.916H4zm1-3h8v.916H5z"/></g>`),
		g.Group(children),
	)
}
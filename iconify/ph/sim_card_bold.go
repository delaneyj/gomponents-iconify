package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SimCardBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m216.49 79.51l-56-56A12 12 0 0 0 152 20H56a20 20 0 0 0-20 20v176a20 20 0 0 0 20 20h144a20 20 0 0 0 20-20V88a12 12 0 0 0-3.51-8.49ZM196 212H60V44h87l49 49ZM88 112a12 12 0 0 0-12 12v64a12 12 0 0 0 12 12h80a12 12 0 0 0 12-12v-64a12 12 0 0 0-12-12Zm12 24h16v40h-16Zm56 40h-16v-40h16Z"/>`),
		g.Group(children),
	)
}
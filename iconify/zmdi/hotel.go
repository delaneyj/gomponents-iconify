package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hotel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M128 213q-27 0-45.5-18.5t-18.5-45T82.5 104T128 85t45.5 19t18.5 45.5t-18.5 45T128 213zM384 85q35 0 60 25t25 61v192h-42v-64H43v64H0V43h43v192h170V85h171z"/>`),
		g.Group(children),
	)
}
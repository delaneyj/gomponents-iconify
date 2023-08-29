package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AirlineSeatIndividualSuite(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M128 213q-26 0-45-18.5t-19-45T83 104t45-19t45 19t19 45.5t-19 45t-45 18.5zM384 85q35 0 60 25t25 61v128H0V85h43v150h170V85h171z"/>`),
		g.Group(children),
	)
}
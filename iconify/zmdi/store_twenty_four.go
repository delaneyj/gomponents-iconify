package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StoreTwentyFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M363 85h64v278H256v-86h-85v86H0V85h64V21h299v64zm-171 64V85h-64v22h43v21h-43v64h64v-21h-43v-22h43zm107 43V85h-22v43h-21V85h-21v64h42v43h22z"/>`),
		g.Group(children),
	)
}
package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ruby(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m192 1024l1-1q-88-4-138-54Q5 920 1 832H0V512h1q0-72 46.5-162T177 177T350 47.5T512 1V0h320v1q88 4 137 54q50 50 54 138l1-1v832H192zm768-64L736 736L448 960h512zM675 713L320 576l-54 376q95-19 202.5-81T675 713zM64 864q2 5 8.5 19.5t12.5 27T96 928q29 29 53 33.5t77-3.5Q68 635 64 626v238zM548.5 91.5q-47.5-47.5-148-14t-195 128t-128 195t14 148t148 14t195-128t128-195t-14-148zM713 675q95-99 157-206.5T951 266l-375 54zM926 96q-5-5-17-11t-27-12.5t-19-8.5H626l330 160q8-52 3.5-76T926 96z"/>`),
		g.Group(children),
	)
}
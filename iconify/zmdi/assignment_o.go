package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AssignmentO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M341 43q18 0 30.5 12.5T384 85v342q0 17-12.5 29.5T341 469H43q-18 0-30.5-12.5T0 427V85q0-17 12.5-29.5T43 43h89q7-19 23.5-31T192 0t36.5 12T252 43h89zm-149 0q-9 0-15 6t-6 15t6 15t15 6t15-6t6-15t-6-15t-15-6zm149 384V85h-42v64H85V85H43v342h298z"/>`),
		g.Group(children),
	)
}
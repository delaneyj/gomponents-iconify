package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Devices(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M85 64v235h214v64H0v-64h43V64q0-18 12.5-30.5T85 21h384v43H85zm406 43q8 0 14.5 6t6.5 15v213q0 9-6.5 15.5T491 363H363q-9 0-15.5-6.5T341 341V128q0-9 6.5-15t15.5-6h128zm-22 192V149h-85v150h85z"/>`),
		g.Group(children),
	)
}
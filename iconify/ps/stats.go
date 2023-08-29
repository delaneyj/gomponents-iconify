package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Stats(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M21 512h86q9 0 15-6t6-15V341q0-17-12.5-29.5T85 299H43q-18 0-30.5 12.5T0 341v150q0 9 6 15t15 6zm22-171h42v43H43v-43zm0 86h42v42H43v-42zm170 85h86q9 0 15-6t6-15V192q0-17-12.5-30T277 149h-42q-18 0-30.5 13T192 192v299q0 9 6 15t15 6zm22-320h42v64h-42v-64zm0 107h42v170h-42V299zM469 0h-42q-18 0-30.5 12.5T384 43v448q0 9 6 15t15 6h86q9 0 15-6t6-15V43q0-18-12.5-30.5T469 0zm0 469h-42v-85h42v85zm0-128h-42V43h42v298z"/>`),
		g.Group(children),
	)
}
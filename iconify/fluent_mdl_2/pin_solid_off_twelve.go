package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PinSolidOffTwelve(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 512v853h-85q-38 0-75-8t-73-21L825 345q78 8 150 35t134 71t113 103t84 129h317q60-81 150-126t190-45h85zM25 146L146 25l1877 1877l-121 121l-689-689q-42 48-93 85t-108 64t-118 39t-126 14h-85v-512H171L0 939l171-86h512v-50L25 146z"/>`),
		g.Group(children),
	)
}
package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func History(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256.25 16A240 240 0 0 0 88 84.977V16H56v128h128v-32h-77.713A208 208 0 0 1 256.25 48C370.8 48 464 141.2 464 255.75S370.8 463.5 256.25 463.5S48.5 370.3 48.5 255.75h-32a239.75 239.75 0 0 0 409.279 169.529A239.75 239.75 0 0 0 256.25 16Z"/><path fill="currentColor" d="M240 111.951L239.465 288H368v-32h-96.437L272 112.049l-32-.098z"/>`),
		g.Group(children),
	)
}
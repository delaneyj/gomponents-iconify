package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThumbDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M299 0q17 0 29.5 12.5T341 43v213q0 18-12 30L188 427l-22-23q-10-9-10-22l1-7l20-98H43q-18 0-30.5-12.5T0 235v-43q0-8 3-16L67 26Q78 0 107 0h192zm85 0h85v256h-85V0z"/>`),
		g.Group(children),
	)
}
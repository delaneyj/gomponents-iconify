package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderStyle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M256 384v-43h43v43h-43zm85 0v-43h43v43h-43zm-256 0v-43h43v43H85zm86 0v-43h42v43h-42zm170-85v-43h43v43h-43zm0-86v-42h43v42h-43zM0 0h384v43H43v341H0V0zm341 128V85h43v43h-43z"/>`),
		g.Group(children),
	)
}
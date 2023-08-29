package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderTop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M85 384v-43h43v43H85zm0-171v-42h43v42H85zm86 0v-42h42v42h-42zm0 171v-43h42v43h-42zM0 299v-43h43v43H0zm0 85v-43h43v43H0zm0-171v-42h43v42H0zm0-85V85h43v43H0zm171 171v-43h42v43h-42zm170-171V85h43v43h-43zm0 85v-42h43v42h-43zM0 0h384v43H0V0zm341 299v-43h43v43h-43zm-85 85v-43h43v43h-43zm-85-256V85h42v43h-42zm170 256v-43h43v43h-43zm-85-171v-42h43v42h-43z"/>`),
		g.Group(children),
	)
}
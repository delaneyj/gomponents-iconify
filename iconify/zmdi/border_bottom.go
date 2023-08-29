package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderBottom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M128 171v42H85v-42h43zm85 85v43h-42v-43h42zM128 0v43H85V0h43zm85 171v42h-42v-42h42zM43 0v43H0V0h43zm170 85v43h-42V85h42zm86 86v42h-43v-42h43zM213 0v43h-42V0h42zm86 0v43h-43V0h43zm42 213v-42h43v42h-43zm0 86v-43h43v43h-43zM43 85v43H0V85h43zM341 0h43v43h-43V0zm0 128V85h43v43h-43zM43 171v42H0v-42h43zM0 384v-43h384v43H0zm43-128v43H0v-43h43z"/>`),
		g.Group(children),
	)
}
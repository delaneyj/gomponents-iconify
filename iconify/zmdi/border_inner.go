package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderInner(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M0 384v-43h43v43H0zm85 0v-43h43v43H85zM43 85v43H0V85h43zM0 299v-43h43v43H0zM128 0v43H85V0h43zM43 0v43H0V0h43zm256 0v43h-43V0h43zm42 128V85h43v43h-43zm0-128h43v43h-43V0zm-85 384v-43h43v43h-43zM213 0v171h171v42H213v171h-42V213H0v-42h171V0h42zm128 384v-43h43v43h-43zm0-85v-43h43v43h-43z"/>`),
		g.Group(children),
	)
}
package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlipToFront(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M0 213v-42h43v42H0zm0 86v-43h43v43H0zm43 85q-18 0-30.5-12.5T0 341h43v43zM0 128V85h43v43H0zm256 256v-43h43v43h-43zM341 0q18 0 30.5 12.5T384 43v213q0 18-12.5 30.5T341 299H128q-18 0-30.5-12.5T85 256V43q0-18 12.5-30.5T128 0h213zm0 256V43H128v213h213zM171 384v-43h42v43h-42zm-86 0v-43h43v43H85z"/>`),
		g.Group(children),
	)
}
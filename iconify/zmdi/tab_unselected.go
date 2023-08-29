package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TabUnselected(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M0 128V85h43v43H0zm0 85v-42h43v42H0zM0 43q0-18 12.5-30.5T43 0v43H0zm171 341v-43h42v43h-42zM0 299v-43h43v43H0zm43 85q-18 0-30.5-12.5T0 341h43v43zM427 0q17 0 29.5 12.5T469 43v85H256V0h171zm0 299v-43h42v43h-42zM171 43V0h42v43h-42zM85 384v-43h43v43H85zm0-341V0h43v43H85zm342 341v-43h42q0 18-12.5 30.5T427 384zm0-171v-42h42v42h-42zM256 384v-43h43v43h-43zm85 0v-43h43v43h-43z"/>`),
		g.Group(children),
	)
}
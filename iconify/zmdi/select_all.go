package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SelectAll(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M0 43q0-18 12.5-30.5T43 0v43H0zm0 170v-42h43v42H0zm85 171v-43h43v43H85zM0 128V85h43v43H0zM213 0v43h-42V0h42zm128 0q18 0 30.5 12.5T384 43h-43V0zM43 384q-18 0-30.5-12.5T0 341h43v43zM0 299v-43h43v43H0zM128 0v43H85V0h43zm43 384v-43h42v43h-42zm170-171v-42h43v42h-43zm0 171v-43h43q0 18-12.5 30.5T341 384zm0-256V85h43v43h-43zm0 171v-43h43v43h-43zm-85 85v-43h43v43h-43zm0-341V0h43v43h-43zM85 299V85h214v214H85zm43-171v128h128V128H128z"/>`),
		g.Group(children),
	)
}
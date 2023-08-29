package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gradient(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M171 128h42v43h-42v-43zm-43 43h43v42h-43v-42zm85 0h43v42h-43v-42zm43-43h43v43h-43v-43zm-171 0h43v43H85v-43zM341 0q18 0 30.5 12.5T384 43v298q0 18-12.5 30.5T341 384H43q-18 0-30.5-12.5T0 341V43q0-18 12.5-30.5T43 0h298zM128 320v-43H85v43h43zm85 0v-43h-42v43h42zm86 0v-43h-43v43h43zm42-149V43H43v128h42v42H43v43h42v-43h43v43h43v-43h42v43h43v-43h43v43h42v-43h-42v-42h42z"/>`),
		g.Group(children),
	)
}
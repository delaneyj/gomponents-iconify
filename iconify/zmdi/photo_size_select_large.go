package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhotoSizeSelectLarge(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M427 256h42v43h-42v-43zm0-85h42v42h-42v-42zm42 170q0 16-13 29.5T427 384v-43h42zM256 0h43v43h-43V0zm171 85h42v43h-42V85zm0-85q16 0 29 13.5T469 43h-42V0zM0 85h43v43H0V85zM341 0h43v43h-43V0zm0 341h43v43h-43v-43zM43 0v43H0q0-16 13.5-29.5T43 0zm128 0h42v43h-42V0zM85 0h43v43H85V0zM0 171h299v213H43q-18 0-30.5-12.5T0 341V171zm43 170h213l-68-91l-54 69l-38-46z"/>`),
		g.Group(children),
	)
}
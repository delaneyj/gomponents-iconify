package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhotoSizeSelectSmall(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M469 256v43h-42v-43h42zm0-85v42h-42v-42h42zm0 170q0 16-13 29.5T427 384v-43h42zM299 0v43h-43V0h43zm170 85v43h-42V85h42zM427 0q16 0 29 13.5T469 43h-42V0zM43 384q-18 0-30.5-12.5T0 341v-85h213v128H43zm0-299v43H0V85h43zm256 256v43h-43v-43h43zM384 0v43h-43V0h43zm0 341v43h-43v-43h43zM43 0v43H0q0-16 13.5-29.5T43 0zm0 171v42H0v-42h43zM213 0v43h-42V0h42zm-85 0v43H85V0h43z"/>`),
		g.Group(children),
	)
}
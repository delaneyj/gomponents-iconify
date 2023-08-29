package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M0 384v-43h43v43H0zM43 85v43H0V85h43zM0 299v-43h43v43H0zm85 85v-43h43v43H85zM43 0v43H0V0h43zm85 0v43H85V0h43zm171 0v43h-43V0h43zm-86 85v43h-42V85h42zm0-85v43h-42V0h42zm128 299v-43h43v43h-43zm-170 85v-43h42v43h-42zM0 213v-42h384v42H0zM341 0h43v43h-43V0zm0 128V85h43v43h-43zM171 299v-43h42v43h-42zm85 85v-43h43v43h-43zm85 0v-43h43v43h-43z"/>`),
		g.Group(children),
	)
}
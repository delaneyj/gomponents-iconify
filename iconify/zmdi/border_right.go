package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M85 384v-43h43v43H85zM0 43V0h43v43H0zm85 0V0h43v43H85zm0 170v-42h43v42H85zM0 384v-43h43v43H0zm171 0v-43h42v43h-42zM0 213v-42h43v42H0zm0 86v-43h43v43H0zm0-171V85h43v43H0zm171 171v-43h42v43h-42zm85-86v-42h43v42h-43zM341 0h43v384h-43V0zm-85 384v-43h43v43h-43zm0-341V0h43v43h-43zm-85 170v-42h42v42h-42zm0-170V0h42v43h-42zm0 85V85h42v43h-42z"/>`),
		g.Group(children),
	)
}
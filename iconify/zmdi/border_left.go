package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M171 384v-43h42v43h-42zm0-85v-43h42v43h-42zm0-256V0h42v43h-42zm0 85V85h42v43h-42zm0 85v-42h42v42h-42zM85 384v-43h43v43H85zm0-341V0h43v43H85zm0 170v-42h43v42H85zM0 384V0h43v384H0zm341-256V85h43v43h-43zm-85 256v-43h43v43h-43zm85-85v-43h43v43h-43zm0-299h43v43h-43V0zm0 213v-42h43v42h-43zm0 171v-43h43v43h-43zm-85-171v-42h43v42h-43zm0-170V0h43v43h-43z"/>`),
		g.Group(children),
	)
}
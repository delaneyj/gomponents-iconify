package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderVertical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M0 128V85h43v43H0zm0-85V0h43v43H0zm85 341v-43h43v43H85zm0-171v-42h43v42H85zm-85 0v-42h43v42H0zm0 171v-43h43v43H0zm0-85v-43h43v43H0zM85 43V0h43v43H85zm256 256v-43h43v43h-43zm-170 85V0h42v384h-42zm170 0v-43h43v43h-43zm0-171v-42h43v42h-43zm0-213h43v43h-43V0zm0 128V85h43v43h-43zm-85-85V0h43v43h-43zm0 341v-43h43v43h-43zm0-171v-42h43v42h-43z"/>`),
		g.Group(children),
	)
}
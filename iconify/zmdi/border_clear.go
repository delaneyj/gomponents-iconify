package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderClear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M85 43V0h43v43H85zm0 170v-42h43v42H85zm0 171v-43h43v43H85zm86-85v-43h42v43h-42zm0 85v-43h42v43h-42zM0 384v-43h43v43H0zm0-85v-43h43v43H0zm0-86v-42h43v42H0zm0-85V85h43v43H0zm0-85V0h43v43H0zm171 170v-42h42v42h-42zm170 86v-43h43v43h-43zm0-86v-42h43v42h-43zm0 171v-43h43v43h-43zm0-256V85h43v43h-43zm-170 0V85h42v43h-42zM341 0h43v43h-43V0zM171 43V0h42v43h-42zm85 341v-43h43v43h-43zm0-171v-42h43v42h-43zm0-170V0h43v43h-43z"/>`),
		g.Group(children),
	)
}
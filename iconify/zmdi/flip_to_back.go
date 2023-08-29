package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlipToBack(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M128 85v43H85V85h43zm0 86v42H85v-42h43zm0-171v43H85q0-18 12.5-30.5T128 0zm85 256v43h-42v-43h42zM341 0q18 0 30.5 12.5T384 43h-43V0zM213 0v43h-42V0h42zm-85 299q-18 0-30.5-12.5T85 256h43v43zm213-86v-42h43v42h-43zm0-85V85h43v43h-43zm0 171v-43h43q0 18-12.5 30.5T341 299zM43 85v256h256v43H43q-18 0-30.5-12.5T0 341V85h43zm213-42V0h43v43h-43zm0 256v-43h43v43h-43z"/>`),
		g.Group(children),
	)
}
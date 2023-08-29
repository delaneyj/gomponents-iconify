package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Memory(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M256 128v128H128V128h128zm-43 85v-42h-42v42h42zm171-42h-43v42h43v43h-43v43q0 17-12.5 29.5T299 341h-43v43h-43v-43h-42v43h-43v-43H85q-17 0-29.5-12.5T43 299v-43H0v-43h43v-42H0v-43h43V85q0-17 12.5-29.5T85 43h43V0h43v43h42V0h43v43h43q17 0 29.5 12.5T341 85v43h43v43zm-85 128V85H85v214h214z"/>`),
		g.Group(children),
	)
}
package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PopMobiili(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.667 27.962v-8.046h2.662a2.716 2.716 0 1 1 0 5.431h-2.662m15.239 2.746v-8.046h2.662a2.716 2.716 0 1 1 0 5.43h-2.662m-7.882-.05a2.857 2.857 0 0 0 5.715 0v-2.854a2.87 2.87 0 0 0-2.912-2.853a2.78 2.78 0 0 0-2.803 2.853Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}
package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberEight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.806 24a6.75 6.75 0 0 0-6.75 6.75h0a6.75 6.75 0 0 0 6.75 6.75h4.388a6.75 6.75 0 0 0 6.75-6.75h0a6.75 6.75 0 0 0-6.75-6.75m0 0a6.75 6.75 0 0 0 6.75-6.75h0a6.75 6.75 0 0 0-6.75-6.75h-4.388a6.75 6.75 0 0 0-6.75 6.75h0a6.75 6.75 0 0 0 6.75 6.75m0 0h4.388"/>`),
		g.Group(children),
	)
}
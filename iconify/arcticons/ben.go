package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ben(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 42.5h-33a2 2 0 0 1-2-2v-33a2 2 0 0 1 2-2h33a2 2 0 0 1 2 2v33a2 2 0 0 1-2 2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.793 28.86a3.249 3.249 0 0 1-2.824 1.64h0a3.25 3.25 0 0 1-3.25-3.25v-2.113a3.25 3.25 0 0 1 3.25-3.25h0a3.25 3.25 0 0 1 3.25 3.25v1.057h-6.5M36.907 30.5v-5.363a3.25 3.25 0 0 0-3.25-3.25h0a3.25 3.25 0 0 0-3.25 3.25V30.5m0-5.363v-3.25"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width=".866" d="M16.456 24a3.25 3.25 0 0 1 0 6.5h-5.363v-13h5.362a3.25 3.25 0 0 1 0 6.5Zm0 0h-5.363"/>`),
		g.Group(children),
	)
}
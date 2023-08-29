package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Symboleuro(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.48 34.27a9 9 0 0 1-6.93 3.23h0a9 9 0 0 1-9-9v-8.9a9 9 0 0 1 9.05-9h0a9 9 0 0 1 7 3.26m-20.43 6.86H22.6m-11.43 6.63H22.6"/>`),
		g.Group(children),
	)
}
package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Goldenone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.45 5.5a2 2 0 0 0-1.95 2v33.1a2 2 0 0 0 2 2h33.1a2 2 0 0 0 2-2V7.45a2 2 0 0 0-2-1.95Zm20.53 28h9.5m-9.5-16.41l4.75-2.59m0 0v19"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.23 20.87a6.36 6.36 0 0 0-6.7-6.36a6.61 6.61 0 0 0-6 6.72v5.9a6.36 6.36 0 0 0 6.35 6.37h0a6.37 6.37 0 0 0 6.36-6.37h-6.37"/>`),
		g.Group(children),
	)
}
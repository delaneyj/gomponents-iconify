package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Oishare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.039 35.13a4.58 4.58 0 0 0-4.128-4.15l-.012.012m.736 3.387a1.836 1.836 0 1 0 0 2.597h0a1.837 1.837 0 0 0 0-2.597Zm7.624.48a9.006 9.006 0 0 0-8.105-8.104m15.141-16.269L19.3 19.482a1 1 0 0 0 0 1.414l6.704 6.704m1.793 8.557l8.996-8.996a1 1 0 0 0 0-1.414l-6.704-6.704"/>`),
		g.Group(children),
	)
}
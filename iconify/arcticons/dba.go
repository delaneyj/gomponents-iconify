package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dba(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.4 25.18a3.45 3.45 0 1 0-6.9 0v2.294a3.45 3.45 0 1 0 6.9 0m0 3.528V16.891m3.318 8.335a3.45 3.45 0 1 1 6.9 0v2.293a3.45 3.45 0 1 1-6.9 0m0 3.528V16.936M37.5 27.58a3.45 3.45 0 1 1-6.9 0v-2.292a3.45 3.45 0 1 1 6.9 0m0 5.82v-9.349"/>`),
		g.Group(children),
	)
}
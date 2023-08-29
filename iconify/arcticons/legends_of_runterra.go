package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LegendsOfRunterra(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.588 31.914V12.5h4.67c2.639 0 4.779 2.145 4.779 4.79s-2.14 4.79-4.78 4.79h-4.67m4.67 0l4.669 9.831m-1.886 0h2.745m-11.57 0h2.745M18.215 12.5h1.373m-5.489 23.283h6.668L24 37.5l3.233-1.717h6.668"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.5 42.5h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2Z"/>`),
		g.Group(children),
	)
}
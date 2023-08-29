package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Screwfix(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.166 15.528h-6.617a1.963 1.963 0 0 0-1.728 2.894l6.014 11.156a1.963 1.963 0 0 1-1.727 2.894H9.5M27.37 21.67l.745-3.129m-3.317 13.931l1.116-4.687m.725-3.044h8.779m-6.586-9.213H38.5"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}
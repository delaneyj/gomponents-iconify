package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UberDriver(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 42.5h-33a2 2 0 0 1-2-2v-33a2 2 0 0 1 2-2h33a2 2 0 0 1 2 2v33a2 2 0 0 1-2 2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.786 18.956A2.726 2.726 0 0 1 38.5 16.23m-2.714 0v7.204m-3.179-1.369a2.75 2.75 0 0 1-5.125-1.251a2.702 2.702 0 0 1-.002-.095v-1.763a2.714 2.714 0 1 1 5.429-.011v.893H27.48M9.5 12.565v7.273a3.608 3.608 0 0 0 7.204 0v-7.273m2.796 6.391a2.714 2.714 0 1 1 5.428 0v1.763a2.714 2.714 0 0 1-2.714 2.715h0a2.714 2.714 0 0 1-2.715-2.715h0m0 2.715V12.565M5.5 33.009q8.721.002 17.44 0l-4.031-4.031m4.031 4.031l-4.031 4.031"/>`),
		g.Group(children),
	)
}
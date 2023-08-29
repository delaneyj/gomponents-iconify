package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Zappka(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.5 42.5h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.621 21.35h4l-4 5.3h4m19.609-8v8m0-1.699l3.622-3.604m-2.469 2.457l2.847 2.834M17.124 24.65a2 2 0 0 1-2 2h0a2 2 0 0 1-2-2v-1.3a2 2 0 0 1 2-2h0a2 2 0 0 1 2 2m0 3.3v-5.3m23.256 3.3a2 2 0 0 1-2 2h0a2 2 0 0 1-2-2v-1.3a2 2 0 0 1 2-2h0a2 2 0 0 1 2 2m-.001 3.3v-5.3m-21.177 3.3a2 2 0 0 0 2 2h0a2 2 0 0 0 2-2v-1.3a2 2 0 0 0-2-2h0a2 2 0 0 0-2 2m0-2v8m6.014-4.7a2 2 0 0 0 2 2h0a2 2 0 0 0 2-2v-1.3a2 2 0 0 0-2-2h0a2 2 0 0 0-2 2m0-2v8m-12.085 1.493a14.641 14.641 0 0 0 16.077 0"/><circle cx="9.621" cy="19.6" r=".75" fill="currentColor"/>`),
		g.Group(children),
	)
}
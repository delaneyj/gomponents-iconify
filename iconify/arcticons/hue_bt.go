package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HueBt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.434 36.5a3.715 3.715 0 0 1-3.262 1.92h0a3.85 3.85 0 0 1-3.838-3.839v-2.494a3.85 3.85 0 0 1 3.838-3.838h0a3.85 3.85 0 0 1 3.838 3.838v1.343h-7.676m-10.172-5.181v6.332A3.85 3.85 0 0 0 24 38.42h0a3.85 3.85 0 0 0 3.838-3.838v-6.333m0 6.332v3.838M9.99 23.067v15.352m0-6.332a3.85 3.85 0 0 1 3.838-3.838h0a3.85 3.85 0 0 1 3.838 3.838v6.333m2.039-25.838l8.59 8.084l-4.463 4.463v-16l4.463 3.958l-8.59 8.168"/>`),
		g.Group(children),
	)
}
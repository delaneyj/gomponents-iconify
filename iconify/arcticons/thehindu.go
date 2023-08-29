package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Thehindu(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 42.5h-33a2.006 2.006 0 0 1-2-2v-33a2.006 2.006 0 0 1 2-2h33a2.006 2.006 0 0 1 2 2v33a2.006 2.006 0 0 1-2 2Zm-31.297-27h11.263m-5.631 17v-17m11.637 0v17m11.262-17v17m-11.262-8.532h11.262M13.772 32.5h2.125m9.512 0h2.125m9.138 0h2.125m-13.388-17h2.125m9.138 0h2.125m-18.331 0v1.063M9.203 15.5v1.063"/>`),
		g.Group(children),
	)
}
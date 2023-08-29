package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Radioactive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m13.5 14.6l3 5.19A9 9 0 0 0 21 12h-6a3 3 0 0 1-1.5 2.6m0-5.2l3-5.19a9 9 0 0 0-9 0l3 5.19a3 3 0 0 1 3 0m-3 5.2l-3 5.19A9 9 0 0 1 3 12h6a3 3 0 0 0 1.5 2.6"/>`),
		g.Group(children),
	)
}
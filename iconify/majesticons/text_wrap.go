package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextWrap(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 7h14M5 12h11c1 0 3 .5 3 2.5S17.333 17 16.5 17H12m-7 0h4m3 0l2-2m-2 2l2 2"/>`),
		g.Group(children),
	)
}
package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoxAlignRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13.998 20.003v-16h5a1 1 0 0 1 1 1v14a1 1 0 0 1-1 1h-5zm-5 0h.01m-5.011 0h.011m-.011-5.001h.011m-.011-6h.011m-.011-5h.011m4.99 0h.01"/>`),
		g.Group(children),
	)
}
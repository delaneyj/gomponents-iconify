package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocationDiscount(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12.797 19.595L10 14l-7-3.5a.55.55 0 0 1 0-1L21 3l-3.548 9.826M16 21l5-5m0 5v.01M16 16v.01"/>`),
		g.Group(children),
	)
}
package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PowerSupplyOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M33 34H15L8 17.75L10 9h28l2 8.75L33 34ZM18 4v5m12-5v5m-6 25v10h16v-7.368M18 21h12"/>`),
		g.Group(children),
	)
}
package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PowerSupplyOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTPowerSupplyOne0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M33 34H15L8 17.75L10 9h28l2 8.75L33 34Z"/><path d="M18 4v5m12-5v5m-6 25v10h16v-7.368M18 21h12"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTPowerSupplyOne0)"/>`),
		g.Group(children),
	)
}
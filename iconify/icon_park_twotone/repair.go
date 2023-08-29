package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Repair(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTRepair0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M17 14h12m-6 14V15m-2.857 27H8a2 2 0 0 1-2-2V7a2 2 0 0 1 2-2h32a2 2 0 0 1 2 2v9.717"/><path fill="#555" d="m27 38l10.5-14.5L42 27L31 42h-4v-4Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTRepair0)"/>`),
		g.Group(children),
	)
}
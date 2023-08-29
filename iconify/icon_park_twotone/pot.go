package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pot(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTPot0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="m34 28l10-4"/><path fill="#555" d="M4 28h30l-.439 3.802A7 7 0 0 1 26.607 38H11.393a7 7 0 0 1-6.954-6.198L4 28Z"/><path d="M19 10v10m-8-8v6m16-6v6"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTPot0)"/>`),
		g.Group(children),
	)
}
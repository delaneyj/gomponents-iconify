package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FinancingTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTFinancingTwo0"><g fill="none" stroke="#fff" stroke-width="4"><path d="M12 9.927V7a3 3 0 0 1 3-3h26a3 3 0 0 1 3 3v26a3 3 0 0 1-3 3h-2.983"/><rect width="34" height="34" x="4" y="10" fill="#555" stroke-linejoin="round" rx="3"/><path stroke-linecap="round" stroke-linejoin="round" d="m15 17l6 6l6-6m-13 8h14m-14 6h14m-7-6v11"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTFinancingTwo0)"/>`),
		g.Group(children),
	)
}
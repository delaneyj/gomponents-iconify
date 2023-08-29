package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HistoryQuery(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTHistoryQuery0"><g fill="none" stroke="#fff" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M42 24V9a3 3 0 0 0-3-3H9a3 3 0 0 0-3 3v30a3 3 0 0 0 3 3h15"/><circle cx="32" cy="32" r="6" fill="#555"/><path stroke-linecap="round" stroke-linejoin="round" d="m37 36l5 4M14 16h20m-20 8h8"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTHistoryQuery0)"/>`),
		g.Group(children),
	)
}
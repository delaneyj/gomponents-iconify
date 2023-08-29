package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Excel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTExcel0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-width="4"><path stroke-linejoin="round" d="M8 15V6a2 2 0 0 1 2-2h28a2 2 0 0 1 2 2v36a2 2 0 0 1-2 2H10a2 2 0 0 1-2-2v-9"/><path d="M31 15h3m-6 8h6m-6 8h6"/><path fill="#555" stroke-linejoin="round" d="M4 15h18v18H4z"/><path stroke-linejoin="round" d="m10 21l6 6m0-6l-6 6"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTExcel0)"/>`),
		g.Group(children),
	)
}
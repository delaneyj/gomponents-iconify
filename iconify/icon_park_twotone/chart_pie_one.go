package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartPieOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTChartPieOne0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M5 24c0 10.493 8.507 19 19 19V26a2 2 0 0 1 2-2h17c0-10.493-8.507-19-19-19S5 13.507 5 24Z"/><path d="M30 30h12v12H30z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTChartPieOne0)"/>`),
		g.Group(children),
	)
}
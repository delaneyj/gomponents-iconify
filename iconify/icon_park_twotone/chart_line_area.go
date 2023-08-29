package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartLineArea(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTChartLineArea0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M6 6v36h36"/><path fill="#555" d="m14 34l8-16l10 9L42 6v28H14Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTChartLineArea0)"/>`),
		g.Group(children),
	)
}
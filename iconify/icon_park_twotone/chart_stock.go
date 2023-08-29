package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartStock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTChartStock0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M6 16h8v16H6z"/><path stroke-linecap="round" d="M10 6v10m0 16v10"/><path fill="#555" d="M34 16h8v16h-8z"/><path stroke-linecap="round" d="M38 6v10m0 16v10"/><path fill="#555" d="M20 14h8v16h-8z"/><path stroke-linecap="round" d="M24 4v10m0 16v10"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTChartStock0)"/>`),
		g.Group(children),
	)
}
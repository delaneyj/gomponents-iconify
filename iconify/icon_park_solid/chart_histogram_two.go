package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartHistogramTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSChartHistogramTwo0"><g fill="none"><path fill="#fff" fill-rule="evenodd" d="M4 42h40H4Z" clip-rule="evenodd"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M4 42h40"/><path fill="#fff" stroke="#fff" stroke-linejoin="round" stroke-width="4" d="M8 28h6v14H8zm13-10h6v24h-6zM34 6h6v36h-6z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSChartHistogramTwo0)"/>`),
		g.Group(children),
	)
}
package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartBarStackedNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsChartBarStackedNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M0 0h48v48H0V0Zm8 40V7H6v34a1 1 0 0 0 1 1h34v-2h-2V13a1 1 0 0 0-1-1h-6a1 1 0 0 0-1 1v27h-2V21a1 1 0 0 0-1-1h-6a1 1 0 0 0-1 1v19h-2V25a1 1 0 0 0-1-1h-6a1 1 0 0 0-1 1v15H8Zm29-16h-4V14h4v10Zm-14-2h4v6h-4v-6Zm-6 10h-4v-6h4v6Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsChartBarStackedNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}
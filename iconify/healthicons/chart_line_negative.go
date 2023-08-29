package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartLineNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsChartLineNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM8 7v25.12l5.33-6.751A3 3 0 1 1 19 23.946l6.633 2.21a2.995 2.995 0 0 1 3.41-.97l6.378-7.653a3 3 0 1 1 1.536 1.28l-6.378 7.654A3 3 0 1 1 25 28.054l-6.633-2.21a2.995 2.995 0 0 1-3.577.902L8 35.347V40h33v2H7a1 1 0 0 1-1-1V7h2Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsChartLineNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}
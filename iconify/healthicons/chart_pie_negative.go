package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartPieNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsChartPieNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM24 6v16.845l15.066-8.698C35.851 9.24 30.304 6 24 6ZM8.934 33.853l31.134-17.975A17.924 17.924 0 0 1 42 24c0 9.941-8.059 18-18 18c-6.304 0-11.851-3.24-15.066-8.147ZM24 4C12.954 4 4 12.954 4 24s8.954 20 20 20s20-8.954 20-20S35.046 4 24 4Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsChartPieNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}
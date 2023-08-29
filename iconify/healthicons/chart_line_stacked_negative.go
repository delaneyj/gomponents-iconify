package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartLineStackedNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsChartLineStackedNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM8 29.387v6.851l8.232-9.878a1 1 0 0 1 1.084-.309l11.353 3.785l9.563-11.476l1.536 1.28l-10 12a1 1 0 0 1-1.084.309L17.33 28.164L8 39.362V40h33v2H7a1 1 0 0 1-1-1V7h2v19.413l8.26-9.086a1 1 0 0 1 1.082-.267l10.34 3.76l9.55-11.46l1.536 1.28l-10 12a1 1 0 0 1-1.11.3l-10.37-3.771L8 29.387Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsChartLineStackedNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}
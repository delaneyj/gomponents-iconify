package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThermometerLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="m17.096 10l1.3 1.3m-4.5 1.9l1.3 1.3m-4.51 1.909l1.301 1.3M5.58 20.558l.383-.384a2.014 2.014 0 0 1 1.647-.578l.799.09a3.02 3.02 0 0 0 2.47-.867l8.942-8.943a4.028 4.028 0 1 0-5.696-5.696L5.18 13.122a3.021 3.021 0 0 0-.866 2.47l.089.8a2.014 2.014 0 0 1-.578 1.646l-.384.383a1.51 1.51 0 0 0 2.137 2.137Z"/><path d="m9 15l6.5-6.5"/></g>`),
		g.Group(children),
	)
}
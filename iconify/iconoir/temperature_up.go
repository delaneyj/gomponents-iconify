package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TemperatureUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M4 12a5 5 0 1 0 6 0m-6 0V3h6v9m0-9h2m-2 3h2m-2 3h2"/><path d="M7 14a2 2 0 1 0 0 4a2 2 0 0 0 0-4Zm0 0V6m12 12V6m0 0l2.5 2.5M19 6l-2.5 2.5"/></g>`),
		g.Group(children),
	)
}